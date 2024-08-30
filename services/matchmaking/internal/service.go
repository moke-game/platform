package internal

import (
	"context"
	"io"

	allocation "agones.dev/agones/pkg/allocation/go"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/gstones/moke-kit/3rd/agones/pkg/agonesfx"
	"github.com/gstones/moke-kit/3rd/cloud/pkg/cloudfx"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/siface"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"open-match.dev/open-match/pkg/pb"

	matchmaking "github.com/moke-game/platform/api/gen/matchmaking/api"
	"github.com/moke-game/platform/services/matchmaking/internal/agones"
	"github.com/moke-game/platform/services/matchmaking/pkg/mmfx"
)

type Service struct {
	logger   *zap.Logger
	feClient pb.FrontendServiceClient
	director *Director
}

func (s *Service) Match(request *matchmaking.MatchRequest, server grpc.ServerStreamingServer[matchmaking.MatchResponse]) error {
	// Create an Open Match CreateTicketRequest with Open Match's public package
	sent := &pb.CreateTicketRequest{
		Ticket: &pb.Ticket{
			SearchFields: &pb.SearchFields{
				Tags: []string{"beta-gameplay"},
			},
		},
	}

	if err := s.director.FetchMatch("game", "mode"); err != nil {
		return err
	}

	if ticket, err := s.feClient.CreateTicket(context.Background(), sent); err != nil {
		s.logger.Error("failed to create ticket", zap.Error(err))
		return err
	} else if res, err := s.watchAssignments(context.Background(), ticket.Id); err != nil {
		s.logger.Error("failed to watch assignments", zap.Error(err))
		return err
	} else if err := server.Send(&matchmaking.MatchResponse{
		GameId: res,
	}); err != nil {
		s.logger.Error("failed to send game id", zap.Error(err))
	}
	return nil
}

func (s *Service) watchAssignments(ctx context.Context, ticketId string) (string, error) {
	req := &pb.WatchAssignmentsRequest{
		TicketId: ticketId,
	}
	stream, err := s.feClient.WatchAssignments(ctx, req)
	if err != nil {
		s.logger.Error("failed to watch assignments", zap.Error(err))
		return "", err
	}

	assignment, err := stream.Recv()
	if err == io.EOF {
		s.logger.Debug("stream connection closed")
		return "", nil
	}
	if err != nil {
		s.logger.Error("failed to receive assignment", zap.Error(err))
		return "", err
	}
	s.logger.Info("assignment received", zap.Any("id", assignment))

	// Do something with the assignment

	// Close the stream when done
	if err := stream.CloseSend(); err != nil {
		s.logger.Error("failed to close stream", zap.Error(err))
		return "", err
	}
	return assignment.Assignment.GetConnection(), nil
}

func (s *Service) RegisterWithGrpcServer(server siface.IGrpcServer) error {
	matchmaking.RegisterMatchServiceServer(server.GrpcServer(), s)
	return nil
}

func CreateService(
	logger *zap.Logger,
	omFrontend, omBackend string,
	funcUrl string,
	funcPort int32,
	agonesClient allocation.AllocationServiceClient,
	accCli *globalaccelerator.Client,
	subNet string,
) (*Service, error) {
	conn, err := grpc.NewClient(omFrontend, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	aCli := agones.CreateAgonesAllocator(logger, agonesClient, accCli, subNet)
	d := CreateDirector(logger, omBackend, funcUrl, funcPort, aCli)
	return &Service{
		logger:   logger,
		director: d,
		feClient: pb.NewFrontendServiceClient(conn),
	}, nil
}

var Module = fx.Provide(
	func(
		logger *zap.Logger,
		setting mmfx.MatchmakingSettingParams,
		agonesClient agonesfx.AllocateParams,
		awsConfig cloudfx.AWSConfigParams,
	) (out sfx.GrpcServiceResult, err error) {

		if svc, e := CreateService(
			logger,
			setting.OMFrontendUrl,
			setting.OMBackendUrl,
			setting.OMFuncUrl,
			setting.OMFuncPort,
			agonesClient.AllocateClient,
			globalaccelerator.NewFromConfig(awsConfig.Config),
			setting.AWSVPCSubnets,
		); e != nil {
			err = e
		} else {
			out.GrpcService = svc
		}
		return
	})
