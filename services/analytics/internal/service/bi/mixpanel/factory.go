package mixpanel

import (
	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/analytics/internal/service/bi"
	"github.com/moke-game/platform/services/analytics/internal/service/bi/mixpanel/internal"
)

func NewDataProcessor(
	logger *zap.Logger,
	mq miface.MessageQueue,
	userId,
	ip,
	token string) (processor bi.DataProcessor, err error) {
	p := new(internal.Processor)
	if e := p.Init(logger, mq, userId, ip, token); e != nil {
		err = e
	} else {
		processor = p
	}
	return
}
