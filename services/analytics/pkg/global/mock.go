package global

import (
	"google.golang.org/grpc"

	pb "github.com/moke-game/platform.git/api/gen/analytics"
)

type Noop struct {
	grpc.ClientStream
}

func (n *Noop) Send(event *pb.AnalyticsEvents) error {
	return nil
}

func (n *Noop) CloseAndRecv() (*pb.Nothing, error) {
	return nil, nil
}
