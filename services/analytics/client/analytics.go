package client

import (
	"context"

	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/logging/slogger"
	"github.com/gstones/moke-kit/server/pkg/sfx"

	pb "github.com/moke-game/platform/api/gen/analytics/api"
	"github.com/moke-game/platform/services/analytics/pkg/analyfx"
)

type AnalyticsClient struct {
	cmd *ishell.Cmd

	username string
	client   pb.AnalyticsServiceClient
}

func (ac *AnalyticsClient) initShell() {
	ac.cmd = &ishell.Cmd{
		Name:    "analytics",
		Help:    "analytics interactive shell",
		Aliases: []string{"AS"},
	}
	ac.initSubCmd()
}

func (ac *AnalyticsClient) initSubCmd() {
	ac.cmd.AddCmd(&ishell.Cmd{
		Name:    "send",
		Help:    "send analytics events",
		Aliases: []string{"S"},
		Func:    ac.sendAnalytic,
	})
}

func (ac *AnalyticsClient) sendAnalytic(c *ishell.Context) {
	eventName := slogger.ReadLine(c, "event name: ")
	js := slogger.ReadLine(c, "properties: ")
	userID := slogger.ReadLine(c, "user id: ")
	deliverTo := []string{"Local", "ThinkingData", "ClickHouse", "Mixpanel"}
	selected := c.Checklist(deliverTo, "choose your deliver to bi", []int{0, 2})
	if len(selected) == 0 {
		return
	}
	var chose []string
	for _, i := range selected {
		chose = append(chose, deliverTo[i])
	}
	var events []*pb.Event
	for _, v := range chose {
		events = append(events, &pb.Event{
			Event:      eventName,
			UserId:     userID,
			Properties: []byte(js),
			DeliverTo:  pb.DeliveryType(pb.DeliveryType_value[v]),
		})
	}

	req := &pb.AnalyticsEvents{
		Events: events,
	}
	if resp, err := ac.client.Analytics(context.Background(), req); err != nil {
		slogger.Warn(c, err)
	} else {
		slogger.Infof(c, "send analytic events result: %v", resp)
	}

}

func CreateAnalyticsClient(url string, username string) (*ishell.Cmd, error) {
	if c, err := analyfx.NewAnalyticsClient(url, sfx.SecuritySettingsParams{}); err != nil {
		return nil, err
	} else {
		aClient := &AnalyticsClient{
			cmd:      &ishell.Cmd{},
			username: username,
			client:   c,
		}
		aClient.initShell()
		return aClient.cmd, nil
	}
}
