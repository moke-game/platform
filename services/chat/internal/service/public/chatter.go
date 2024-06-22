package public

import (
	"fmt"
	"time"

	"github.com/gstones/moke-kit/mq/common"
	"github.com/gstones/moke-kit/mq/miface"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/gstones/platform/api/gen/chat"
	"github.com/gstones/platform/services/chat/internal/service/db"
)

type Chatter struct {
	logger      *zap.Logger
	deployment  string
	appId       string
	uid         string
	conn        pb.ChatService_ChatServer
	mq          miface.MessageQueue
	latChatTime time.Time
	intervalMax time.Duration
	db          *db.Database

	subscripts map[string]miface.Subscription
}

func CreateChatter(
	uid string,
	deployment string,
	appId string,
	conn pb.ChatService_ChatServer,
	logger *zap.Logger,
	mq miface.MessageQueue,
	intervalMax time.Duration,
	db *db.Database,
) *Chatter {
	return &Chatter{
		logger:      logger,
		uid:         uid,
		conn:        conn,
		mq:          mq,
		deployment:  deployment,
		appId:       appId,
		intervalMax: intervalMax,
		db:          db,
	}
}

func (c *Chatter) Init() {
	c.subscripts = make(map[string]miface.Subscription)
}

func (c *Chatter) Update() {
	defer c.Destroy()
	for {
		select {
		case <-c.conn.Context().Done():
			c.logger.Info("connection closed")
			return
		default:
			in, err := c.conn.Recv()
			if err != nil {
				if status.Code(err) == codes.Canceled {
					c.logger.Info("connection closed")
					return
				}
				c.logger.Error("failed to receive message", zap.Error(err))
				return
			}
			c.handleMsg(in)
		}
	}
}

func (c *Chatter) handleMsg(msg *pb.ChatRequest) {
	switch msg.Kind.(type) {
	case *pb.ChatRequest_Subscribe_:
		c.subscribe(msg.GetSubscribe())
	case *pb.ChatRequest_Unsubscribe:
		c.unSubscribe(msg.GetUnsubscribe())
	case *pb.ChatRequest_Message:
		c.pubMessage(msg.GetMessage())
	}
}

func (c *Chatter) Destroy() {
	c.subscripts = nil
}

func (c *Chatter) unSubscribe(msg *pb.ChatRequest_UnSubscribe) {
	if msg == nil || msg.GetDestination() == nil {
		return
	}
	topic := c.makeChatTopic(msg.GetDestination().GetChannel(), msg.GetDestination().GetId())
	if sub, ok := c.subscripts[topic]; ok {
		if err := sub.Unsubscribe(); err != nil {
			c.logger.Error("failed to unsubscribe", zap.Error(err))
			return
		}
		delete(c.subscripts, topic)
	}
}

func (c *Chatter) subscribe(msg *pb.ChatRequest_Subscribe) {
	if msg == nil || msg.GetDestination() == nil {
		return
	}
	topic := c.makeChatTopic(msg.GetDestination().GetChannel(), msg.GetDestination().GetId())
	if _, ok := c.subscripts[topic]; ok {
		c.logger.Warn("already subscribed", zap.String("topic", topic))
		return
	}
	if sub, err := c.mq.Subscribe(c.conn.Context(), topic, func(msg miface.Message, err error) common.ConsumptionCode {
		if err != nil {
			c.logger.Error("failed to consume message", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		pMsg := &pb.ChatMessage{}
		if err := proto.Unmarshal(msg.Data(), pMsg); err != nil {
			c.logger.Error("failed to unmarshal message", zap.Error(err))
			return common.ConsumeNackPersistentFailure
		}
		if err := c.conn.Send(&pb.ChatResponse{
			Kind: &pb.ChatResponse_Message{
				Message: pMsg,
			},
		}); err != nil {
			c.logger.Error("failed to send message", zap.Error(err))
		}
		return common.ConsumeAck
	}); err != nil {
		c.logger.Error("failed to subscribe", zap.Error(err))
	} else {
		c.subscripts[topic] = sub
	}
}

func (c *Chatter) makeChatTopic(channel int32, id string) string {
	topic := fmt.Sprintf("%s.%s.%d", c.appId, c.deployment, channel)
	if id != "" && id != "0" {
		topic = fmt.Sprintf("%s.%s", topic, id)
	}
	return common.NatsHeader.CreateTopic(topic)
}

func (c *Chatter) sendResponseErr(code pb.ChatError_Code) {
	if code == pb.ChatError_CODE_NONE {
		return
	}
	resp := &pb.ChatResponse{
		Kind: &pb.ChatResponse_Error{
			Error: &pb.ChatError{
				Code: code,
			},
		},
	}
	if err := c.conn.Send(resp); err != nil {
		c.logger.Error("failed to send message", zap.Error(err))
	}
}

func (c *Chatter) pubMessage(message *pb.ChatMessage) {
	if isBlocked, err := c.db.IsBlocked(c.uid); err != nil {
		c.logger.Error("failed to check blocked", zap.Error(err))
		return
	} else if isBlocked {
		c.sendResponseErr(pb.ChatError_CODE_BLOCKED)
		return
	}
	if message == nil || message.GetDestination() == nil {
		return
	}
	if time.Now().Sub(c.latChatTime) < c.intervalMax {
		c.logger.Warn("chat interval too short")
		c.sendResponseErr(pb.ChatError_CODE_INTERVAL)
		return
	}
	c.latChatTime = time.Now()
	topic := c.makeChatTopic(message.GetDestination().GetChannel(), message.GetDestination().GetId())
	if data, err := proto.Marshal(message); err != nil {
		c.logger.Error("failed to marshal message", zap.Error(err))
		return
	} else if err := c.mq.Publish(topic, miface.WithBytes(data)); err != nil {
		c.logger.Error("failed to publish message", zap.Error(err))
	}
}
