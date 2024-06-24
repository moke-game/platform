package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"

	"github.com/moke-game/platform/services/analytics/internal/service/bi"
)

type Processor struct {
	logger   *zap.Logger
	rootPath string
	hostname string
}

func (p *Processor) Init(
	logger *zap.Logger,
	hostname, rootPath string,
) error {
	p.logger = logger.With(zap.String("data_processor", "local_bi"))
	p.hostname = hostname
	p.rootPath = rootPath
	return nil
}

func (p *Processor) Handle(event bi.EventType, properties []byte) error {
	if err := p.deliver(event.String(), properties); err != nil {
		return err
	}
	return nil
}

func (p *Processor) makeLogPath(eventName string) string {
	now := time.Now()
	t := now.Format("2006-01-02")
	h := now.Hour()
	return filepath.Join(p.rootPath, fmt.Sprintf("mta-%s_%s_%d.log", eventName, t, h))
}

func (p *Processor) deliver(eventName string, data []byte) error {
	p.logger.Info("deliver local file bi logs", zap.String("data", string(data)))
	fileName := p.makeLogPath(eventName)
	if err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		return err
	}
	if _, err := file.WriteString("\n"); err != nil {
		return err
	}
	return nil
}
