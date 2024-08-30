package agones

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	allocation "agones.dev/agones/pkg/allocation/go"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	namespace        = "default"
	Battle    string = "battle"
	World     string = "world"
)

type PortMap map[int32][]types.SocketAddress

type Allocator struct {
	logger *zap.Logger
	client allocation.AllocationServiceClient
	accCli *globalaccelerator.Client
	subNet string

	accPortCache *sync.Map
}

func CreateAgonesAllocator(
	logger *zap.Logger,
	client allocation.AllocationServiceClient,
	accCli *globalaccelerator.Client,
	subNet string,
) *Allocator {
	return &Allocator{
		logger:       logger,
		client:       client,
		accCli:       accCli,
		subNet:       subNet,
		accPortCache: new(sync.Map),
	}
}

//func (a *Allocator) Allocate(playId int32, uids ...string) (string, error) {
//	if playId == 0 {
//		if url, err := a.allocateWorld(uids...); err != nil {
//			return "", err
//		} else {
//			return url, nil
//		}
//	} else if url, err := a.AllocateBattle(); err != nil {
//		return "", err
//	} else {
//		return url, nil
//	}
//}
//
//func (a *Allocator) allocateWorld(uids ...string) (string, error) {
//	req, err := a.makeWorldRequest(uids...)
//	if err != nil {
//		return "", err
//	}
//	return a.allocate(req)
//}

func (a *Allocator) AllocateBattle() (string, error) {
	req, err := a.makeBattleRequest()
	if err != nil {
		return "", err
	}
	return a.allocate(req)
}

// world use list feature
func (a *Allocator) makeWorldRequest(uids ...string) (*allocation.AllocationRequest, error) {
	if len(uids) <= 0 {
		return nil, fmt.Errorf("uids is empty")
	}
	matchLabels := map[string]string{
		"agones.dev/fleet": World,
	}
	listSelect := map[string]*allocation.ListSelector{
		"players": {
			//ContainsValue: target,
			MinAvailable: int64(len(uids)),
		},
	}
	selectors := []*allocation.GameServerSelector{
		{
			MatchLabels:     matchLabels,
			GameServerState: allocation.GameServerSelector_ALLOCATED,
			Lists:           listSelect,
		},
		{
			MatchLabels:     matchLabels,
			GameServerState: allocation.GameServerSelector_READY,
		},
	}

	return &allocation.AllocationRequest{
		Namespace: namespace,
		Metadata: &allocation.MetaPatch{
			Labels: map[string]string{"agones.dev/sdk-player-add": uids[0]},
		},
		GameServerSelectors: selectors,
		Lists: map[string]*allocation.ListAction{
			"players": {
				AddValues: uids,
			},
		},
	}, nil
}

// battle use counter feature
func (a *Allocator) makeBattleRequest() (*allocation.AllocationRequest, error) {
	matchLabels := map[string]string{
		"agones.dev/fleet": Battle,
	}
	countersSelect := map[string]*allocation.CounterSelector{
		"rooms": {
			MinAvailable: 1,
		},
	}
	selectors := []*allocation.GameServerSelector{
		{
			GameServerState: allocation.GameServerSelector_ALLOCATED,
			MatchLabels:     matchLabels,
			Counters:        countersSelect,
		},
		{
			MatchLabels:     matchLabels,
			GameServerState: allocation.GameServerSelector_READY,
		},
	}
	countersActions := map[string]*allocation.CounterAction{
		"rooms": {
			Action: &wrapperspb.StringValue{Value: "Increment"},
			Amount: &wrapperspb.Int64Value{Value: 1},
		},
	}

	return &allocation.AllocationRequest{
		Namespace:           namespace,
		GameServerSelectors: selectors,
		Counters:            countersActions,
	}, nil
}

func (a *Allocator) allocate(req *allocation.AllocationRequest) (string, error) {
	if resp, err := a.client.Allocate(context.Background(), req); err != nil {
		return "", err
	} else if len(resp.Ports) <= 0 {
		return "", fmt.Errorf("agones allocated service has no port")
	} else {
		host := fmt.Sprintf("%s:%d", resp.Address, resp.Ports[0].Port)
		internalIp, port := a.getInternalIp(resp)
		if internalIp == "" {
			a.logger.Warn("Agones allocated service has no internalIp", zap.Any("response", resp))
			return host, nil
		}
		accEnds := a.getFromCache(internalIp, port)
		if accEnds == nil || len(accEnds) <= 0 {
			if err := a.intAccCacheFromAws(internalIp); err != nil {
				a.logger.Error("Failed to get GlobalAccelerator endpoints", zap.Error(err))
				return host, nil
			}
			if accEnds = a.getFromCache(internalIp, port); accEnds == nil || len(accEnds) <= 0 {
				a.logger.Warn("GlobalAccelerator has no endpoint", zap.String("internalIp", internalIp), zap.Int32("port", port))
				return host, nil
			}
		}
		index := rand.Intn(len(accEnds))
		host = fmt.Sprintf("%s:%d", *accEnds[index].IpAddress, *accEnds[index].Port)
		a.logger.Info("Agones allocated service", zap.Any("response", resp), zap.String("host", host))
		return host, nil
	}
}

func (a *Allocator) getInternalIp(msg *allocation.AllocationResponse) (string, int32) {
	if msg == nil {
		return "", 0
	}

	port := msg.GetPorts()[0].GetPort()
	for _, v := range msg.GetAddresses() {
		if v.GetType() == "InternalIP" {
			return v.GetAddress(), port
		}
	}
	return "", 0
}

func (a *Allocator) setAccPortCache(internalIp string, maps PortMap) {
	a.accPortCache.Store(internalIp, maps)
}

func (a *Allocator) getFromCache(internalIp string, port int32) []types.SocketAddress {
	if v, ok := a.accPortCache.Load(internalIp); ok {
		if cache, ok1 := v.(PortMap); ok1 {
			if p, ok2 := cache[port]; ok2 {
				return p
			}
		}
	}
	return nil
}

func (a *Allocator) intAccCacheFromAws(internalIp string) error {
	if internalIp == "" {
		return fmt.Errorf("internalIp is empty")
	}
	var num = int32(1000)
	res, err := a.accCli.ListCustomRoutingPortMappingsByDestination(
		context.Background(),
		&globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput{
			DestinationAddress: &internalIp,
			EndpointId:         &a.subNet,
			MaxResults:         &num,
		})
	if err != nil {
		a.logger.Error("Failed to get GlobalAccelerator endpoints", zap.Error(err))
		return err
	}
	var portMap = map[int32][]types.SocketAddress{}
	for _, v := range res.DestinationPortMappings {
		if v.AcceleratorSocketAddresses == nil || len(v.AcceleratorSocketAddresses) <= 0 {
			continue
		}
		destPort := *v.DestinationSocketAddress.Port
		portMap[destPort] = v.AcceleratorSocketAddresses
	}
	a.setAccPortCache(internalIp, portMap)
	return nil
}
