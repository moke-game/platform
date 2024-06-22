package internal

import (
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/gstones/platform/services/analytics/internal/service/bi"
)

var EventTypeName = map[bi.EventType]string{
	bi.EventTypeUserSet:     "user_set",
	bi.EventTypeUserSetOnce: "user_setOnce",
	bi.EventTypeUserAdd:     "user_add",
	bi.EventTypeUserDel:     "user_del",
}

func CreateEvent(eventType bi.EventType, userID, ip string, properties []byte) (event Event, err error) {
	err = event.Init(eventType, userID, ip, properties)
	return
}

type Event struct {
	UserID     string
	IP         string
	Time       string
	EventType  string
	EventName  string
	Properties string
}

func (e *Event) Init(event bi.EventType, userID, ip string, properties []byte) error {
	eventStr := event.String()
	if strings.HasPrefix(eventStr, "#") {
		tp, ok := EventTypeName[event]
		if !ok {
			err := errors.Wrap(bi.ErrNotFoundEventType, eventStr)
			return err
		}
		e.EventType = tp
	} else {
		e.EventType = "track"
		e.EventName = filterSpace(eventStr)
	}
	e.Properties = string(properties)
	e.UserID = userID
	e.IP = ip
	e.Time = time.Now().Format("2006-01-02 15:04:05")

	return nil
}

func filterSpace(str string) string {
	return strings.ReplaceAll(strings.TrimSpace(str), " ", "_")
}
