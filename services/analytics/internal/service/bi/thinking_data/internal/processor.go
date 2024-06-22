package internal

import (
	"bytes"
	"net/url"

	"github.com/gstones/moke-kit/mq/miface"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/gstones/platform/services/analytics/internal/service/bi"
)

type Processor struct {
	logger *zap.Logger
	mq     miface.MessageQueue
	userId string
	ip     string
	token  string
}

func (p *Processor) Init(logger *zap.Logger, mq miface.MessageQueue, userId, ip, token string) error {
	p.logger = logger.With(zap.String("data_processor", "ThinkingData"))
	p.userId = userId
	p.ip = ip
	p.mq = mq
	p.token = token
	return nil
}

func (p *Processor) Handle(name bi.EventType, properties []byte) error {
	if event, err := CreateEvent(name, p.userId, p.ip, properties); err != nil {
		return err
	} else if data, err := p.pack(event); err != nil {
		return err
	} else if err := p.deliver(data); err != nil {
		return err
	}
	return nil
}

func (p *Processor) pack(event Event) ([]byte, error) {
	var tpl bytes.Buffer
	if err := tdTemplate.Execute(&tpl, event); err != nil {
		return nil, err
	} else if ok := jsoniter.Valid(tpl.Bytes()); !ok {
		return nil, errors.Wrap(bi.ErrInvalidProperties, event.Properties)
	} else {
		data := url.Values{}
		data.Add("appid", p.token)
		data.Add("data", tpl.String())
		return []byte(data.Encode()), nil
	}
}

func (p *Processor) deliver(data []byte) error {
	// TODO deliver bi logs
	p.logger.Info("deliver bi logs", zap.String("data", string(data)))
	return nil
}
