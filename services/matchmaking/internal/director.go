package internal

import (
	"context"
	"io"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"open-match.dev/open-match/pkg/pb"

	"github.com/moke-game/platform/services/matchmaking/internal/agones"
)

type Director struct {
	logger     *zap.Logger
	be         pb.BackendServiceClient
	backendUrl string
	funcHost   string
	funcPort   int32
	agonesCli  *agones.Allocator
}

func CreateDirector(
	logger *zap.Logger,
	backendUrl string,
	funcHost string,
	funcPort int32,
	aCli *agones.Allocator,
) *Director {
	return &Director{
		logger:     logger,
		backendUrl: backendUrl,
		funcHost:   funcHost,
		funcPort:   funcPort,
		agonesCli:  aCli,
	}
}

func (d *Director) FetchMatch(game string, mode string) error {
	conn, err := grpc.NewClient(d.backendUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		d.logger.Error("failed to create client", zap.Error(err))
		return err
	}
	client := pb.NewBackendServiceClient(conn)

	req := &pb.FetchMatchesRequest{
		Config: &pb.FunctionConfig{
			Host: d.funcHost,
			Port: d.funcPort,
			Type: pb.FunctionConfig_GRPC,
		},
		Profile: &pb.MatchProfile{
			Name: game,
			Pools: []*pb.Pool{
				{
					Name: mode,
				},
			},
		},
	}

	stream, err := client.FetchMatches(context.Background(), req)
	if err != nil {
		d.logger.Error("failed to fetch matches", zap.Error(err))
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			d.logger.Info("stream closed")
			return nil
		}
		if err != nil {
			d.logger.Error("failed to receive match", zap.Error(err))
			return err
		}
		d.logger.Info("match received", zap.Any("match", resp.GetMatch()))
		d.allocateGameServer(context.Background(), resp.GetMatch())
	}
}

func (d *Director) allocateGameServer(ctx context.Context, match *pb.Match) {
	ids := make([]string, 0)

	for _, t := range match.GetTickets() {
		ids = append(ids, t.GetId())
	}

	gameUrl, err := d.agonesCli.AllocateBattle()
	if err != nil {
		d.logger.Error("failed to allocate game server", zap.Error(err))
		return
	}

	// Allocate a game server
	req := &pb.AssignTicketsRequest{
		Assignments: []*pb.AssignmentGroup{
			{
				TicketIds: ids,
				Assignment: &pb.Assignment{
					Connection: gameUrl,
				},
			},
		},
	}

	resp, err := d.be.AssignTickets(ctx, req)
	if err != nil {
		d.logger.Error("failed to assign tickets", zap.Error(err))
		return
	}
	if len(resp.GetFailures()) > 0 {
		d.logger.Error("failed to assign tickets", zap.Any("failures", resp.GetFailures()))
		return
	}
}
