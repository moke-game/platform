package thinkingdata

import (
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/analytics/internal/service/bi"
	"github.com/moke-game/platform/services/analytics/internal/service/bi/thinkingdata/internal"
)

func NewDataProcessor(
	logger *zap.Logger,
	path string,
) (processor bi.DataProcessor, err error) {
	p := new(internal.Processor)
	if e := p.Init(logger, path); e != nil {
		err = e
	} else {
		processor = p
	}
	return
}
