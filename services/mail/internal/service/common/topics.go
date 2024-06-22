package common

import (
	"fmt"

	"github.com/gstones/moke-kit/mq/common"
)

func MakePrivateTopic(uid string) string {
	return common.NatsHeader.CreateTopic(fmt.Sprintf("mail.private.%s", uid))
}

func MakePublicTopic(channel string) string {
	if channel == "" {
		channel = "0"
	}
	return common.NatsHeader.CreateTopic(fmt.Sprintf("mail.public.%s", channel))
}
