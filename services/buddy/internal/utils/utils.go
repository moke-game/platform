package utils

import (
	"fmt"

	"github.com/gstones/moke-kit/mq/common"
)

func MakeBuddyTopic(uid string) string {
	return common.NatsHeader.CreateTopic(fmt.Sprintf("buddy.changes.%s", uid))
}
