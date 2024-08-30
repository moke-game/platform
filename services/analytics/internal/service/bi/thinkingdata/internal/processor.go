package internal

import (
	"github.com/ThinkingDataAnalytics/go-sdk/v2/src/thinkingdata"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/analytics/internal/service/bi"
)

type Processor struct {
	logger *zap.Logger
	tda    thinkingdata.TDAnalytics
}

func (p *Processor) Init(logger *zap.Logger, path string) error {
	p.logger = logger.With(zap.String("data_processor", "ThinkingData"))
	consumer, err := thinkingdata.NewLogConsumerWithConfig(thinkingdata.TDLogConsumerConfig{
		Directory: path,
	})
	if err != nil {
		return err
	}
	p.tda = thinkingdata.New(consumer)
	return nil
}

func (p *Processor) Handle(event bi.EventType, userID string, distinct string, properties []byte) error {
	return p.handleWithEventType(event, userID, distinct, properties)
}

func (p *Processor) handleWithEventType(event bi.EventType, userID string, distinct string, properties []byte) error {
	proper := map[string]interface{}{}
	if err := jsoniter.Unmarshal(properties, &proper); err != nil {
		return err
	}
	p.logger.Debug(
		"bi data handle",
		zap.String("event", event.String()),
		zap.String("userID", userID),
		zap.String("distinct", distinct),
		zap.Any("properties", proper),
	)
	switch event {
	case bi.EventTypeUserSet:
		return p.tda.UserSet(userID, distinct, proper)
	case bi.EventTypeUserSetOnce:
		return p.tda.UserSetOnce(userID, distinct, proper)
	case bi.EventTypeUserAdd:
		return p.tda.UserAdd(userID, distinct, proper)
	case bi.EventTypeUserDel:
		return p.tda.UserDelete(userID, distinct)
	default:
		return p.tda.Track(userID, distinct, event.String(), proper)
	}
}
