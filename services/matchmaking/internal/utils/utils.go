package utils

import (
	"fmt"

	"github.com/gstones/moke-kit/mq/common"
)

const (
	NTF_MatchingSuccess          = 117
	S2C_EVENT_S2C_MatchingCancel = 46015
	QUEUE_LEVEL_0                = 0
	QUEUE_LEVEL_1                = 1
	QUEUE_LEVEL_2                = 2
	QUEUE_LEVEL_3                = 3
	QUEUE_TIME_1                 = 3
	QUEUE_TIME_2                 = 6
	QUEUE_TIME_3                 = 15
	QUEUE_TIME_4                 = 30
)

func MakeNotifyTopic(uid string) string {
	return common.NatsHeader.CreateTopic(fmt.Sprintf("notification.response.%s", uid))
}
