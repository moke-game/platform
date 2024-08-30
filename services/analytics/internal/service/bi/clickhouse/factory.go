package clickhouse

import (
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/analytics/internal/service/bi"
	"github.com/moke-game/platform/services/analytics/internal/service/bi/clickhouse/internal"
)

func NewDataProcessor(
	logger *zap.Logger, rootPath, addr, dbname, uname, passwd string,
) (processor bi.DataProcessor, err error) {
	p := new(internal.Processor)
	if e := p.Init(logger, rootPath, addr, dbname, uname, passwd); e != nil {
		err = e
	} else {
		processor = p
	}
	return
}
