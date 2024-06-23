package thinking_data

import (
	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"

	"github.com/moke-game/platform.git/services/analytics/internal/service/bi"
	"github.com/moke-game/platform.git/services/analytics/internal/service/bi/thinking_data/internal"
)

func NewDataProcessor(
	logger *zap.Logger,
	mq miface.MessageQueue,
	userId,
	ip,
	token string,
) (processor bi.DataProcessor, err error) {
	p := new(internal.Processor)
	if e := p.Init(logger, mq, userId, ip, token); e != nil {
		err = e
	} else {
		processor = p
	}
	return
}
