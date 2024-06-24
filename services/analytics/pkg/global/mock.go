package global

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/moke-game/platform/api/gen/analytics"
)

type Noop struct {
	grpc.ClientStream
}

func (c *Noop) Analytics(_ context.Context, _ *pb.AnalyticsEvents, _ ...grpc.CallOption) (*pb.Nothing, error) {

	return &pb.Nothing{}, nil
}
