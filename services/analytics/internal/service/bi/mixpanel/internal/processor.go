package internal

import (
	"bytes"

	"github.com/gstones/moke-kit/mq/miface"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/analytics/internal/service/bi"
)

type Processor struct {
	logger *zap.Logger
	mq     miface.MessageQueue
	userId string
	ip     string
	token  string
}

func (p *Processor) Init(logger *zap.Logger, mq miface.MessageQueue, userId, ip, token string) error {
	p.logger = logger.With(zap.String("data_processor", "MixPanel"))
	p.userId = userId
	p.ip = ip
	p.mq = mq
	p.token = token
	return nil
}

func (p *Processor) Handle(name bi.EventType, UserId, distinct string, properties []byte) error {
	if event, err := CreateEvent(name, p.token, p.userId, p.ip, properties); err != nil {
		return err
	} else if data, err := p.pack(event); err != nil {
		return err
	} else if err := p.deliver(event.Topic.String(), data); err != nil {
		return err
	}
	return nil
}

func (p *Processor) pack(event Event) ([]byte, error) {
	if event.Topic == AnalyticsTopicMPUserProfiles {
		return p.packageUserProfiles(event)
	}
	return p.packageTrackEvents(event)
}

func (p *Processor) deliver(topic string, data []byte) error {

	p.logger.Info("deliver bi logs", zap.String("data", string(data)))
	return nil
	//size := base64.StdEncoding.EncodedLen(len(data))
	//payload := make([]byte, size)
	//base64.StdEncoding.Encode(payload, data)
	//
	//
	//return p.mq.Publish(mq.NsqProtocol+topic, mq.WithBytes(payload))
}

func (p *Processor) packageTrackEvents(event Event) ([]byte, error) {
	var tpl bytes.Buffer
	if err := mpTrackTemplate.Execute(&tpl, event); err != nil {
		return nil, err
	} else if ok := jsoniter.Valid(tpl.Bytes()); !ok {
		return nil, errors.Wrap(bi.ErrInvalidProperties, event.Properties)
	} else {
		return tpl.Bytes(), nil
	}
}

func (p *Processor) packageUserProfiles(event Event) ([]byte, error) {
	var tpl bytes.Buffer
	if err := mpEngageTemplate.Execute(&tpl, event); err != nil {
		return nil, err
	} else if ok := jsoniter.Valid(tpl.Bytes()); !ok {
		return nil, errors.Wrap(bi.ErrInvalidProperties, event.Properties)
	}
	return tpl.Bytes(), nil
}
