package local

import (
	"go.uber.org/zap"

	"github.com/moke-game/platform.git/services/analytics/internal/service/bi"
	"github.com/moke-game/platform.git/services/analytics/internal/service/bi/local/internal"
)

func NewDataProcessor(
	logger *zap.Logger,
	hostname string,
	rootPath string,
) (processor bi.DataProcessor, err error) {
	p := new(internal.Processor)
	if e := p.Init(logger, hostname, rootPath); e != nil {
		err = e
	} else {
		processor = p
	}
	return
}
