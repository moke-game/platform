package internal

import (
	"github.com/gstones/moke-kit/mq/common"
)

func makeThinkingDataTopics() string {
	return common.NatsHeader.CreateTopic("analytics.td")
}
