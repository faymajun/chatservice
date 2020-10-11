package service

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net/session"
	"cc-be-chat-test/go_modules/server/chat"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"time"
)

var logger = logrus.WithField("component", "client")
var ClientService = newClientService()

type clientService struct {
	logger logrus.FieldLogger
}

func newClientService() *clientService {
	return &clientService{
		logger: logrus.WithField("service", "client"),
	}
}

// Login client 登录
func (cs *clientService) Login(s *session.Session, data proto.Message) error {
	msg := data.(*message.ReqLogin)
	// 设置名字，ID，登录时间
	s.Set(chat.Name, msg.Name)
	s.Bind(msg.UserId)
	s.Set(chat.LoginTime, time.Now().Unix())

	chat.Room.Login(s)
	logger.Println(msg)
	return nil
}

// Chat 发送聊天信息
func (cs *clientService) Chat(s *session.Session, data proto.Message) error {
	msg := data.(*message.Chat)
	return chat.Room.Chat(s, msg)
}

// Heartbeat 心跳数据
func (cs *clientService) Heartbeat(s *session.Session, data proto.Message) error {
	msg := data.(*message.ReqHeartbeat)
	logger.Println(msg)
	response := &message.ResHeartbeat{Uid: msg.Uid, ServerUnixTime: time.Now().Unix() * 1000}
	return s.Send(message.MSGID_ResHeartbeatE, response)
}
