package changes

import (
	"fmt"

	"github.com/gstones/moke-kit/mq/common"
)

func MakeProfileTopic(uid string) string {
	return common.NatsHeader.CreateTopic(fmt.Sprintf("profile.changes.%s", uid))
}
