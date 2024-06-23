package changes

import (
	"fmt"

	"github.com/gstones/moke-kit/mq/common"
	"google.golang.org/protobuf/proto"

	pb "github.com/moke-game/platform/api/gen/knapsack"
)

func CreateTopic(uid string) string {
	return common.NatsHeader.CreateTopic(fmt.Sprintf("knapsack.changes.%s", uid))
}

func Pack(msg *pb.KnapsackModify) ([]byte, error) {
	return proto.Marshal(msg)
}

func UnPack(data []byte) (*pb.KnapsackModify, error) {
	msg := &pb.KnapsackModify{}
	if err := proto.Unmarshal(data, msg); err != nil {
		return nil, err
	}
	return msg, nil
}
