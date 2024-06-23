package internal

import (
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/moke-game/platform/services/analytics/internal/service/bi"
)

var EventTypeName = map[bi.EventType]string{
	bi.EventTypeUserSet:     "$set",
	bi.EventTypeUserSetOnce: "$set_once",
	bi.EventTypeUserAdd:     "$add",
	bi.EventTypeUserDel:     "$delete",
}

func CreateEvent(eventType bi.EventType, token, userID, ip string, properties []byte) (event Event, err error) {
	err = event.init(eventType, token, userID, ip, properties)
	return
}

type Event struct {
	UserID     string
	Token      string
	IP         string
	Time       string
	Topic      MPTopicType
	EventName  string
	Properties string
}

func (e *Event) init(event bi.EventType, token, userID, ip string, properties []byte) error {
	eventStr := event.String()
	e.UserID = userID
	e.Token = token
	e.IP = ip
	e.Time = time.Now().Format("2006-01-02 15:04:05")
	if strings.HasPrefix(eventStr, "#") {
		tp, ok := EventTypeName[event]
		if !ok {
			err := errors.Wrap(bi.ErrNotFoundEventType, eventStr)
			return err
		}
		e.EventName = tp
		e.Topic = AnalyticsTopicMPUserProfiles
	} else {
		e.Topic = AnalyticsTopicMPTrack
		e.EventName = eventStr
	}

	e.Properties = string(properties)
	return nil
}
