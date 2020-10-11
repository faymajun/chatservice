package service

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net/session"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

var logger = logrus.WithField("component", "client")
var ClientService = newClientService()

// clientService 接受服务消息并相应处理
type clientService struct {
	logger logrus.FieldLogger
}

func newClientService() *clientService {
	return &clientService{
		logger: logrus.WithField("service", "client"),
	}
}

// Login 登录返回
func (cs *clientService) Login(s *session.Session, data proto.Message) error {
	msg := data.(*message.ReqLogin)
	logger.Println(msg)
	return nil
}

// Chat 聊天消息
func (cs *clientService) Chat(s *session.Session, data proto.Message) error {
	msg := data.(*message.Chat)
	logger.Println(msg)
	return nil
}

// History 历时记录
func (cs *clientService) History(s *session.Session, data proto.Message) error {
	msg := data.(*message.HistroyChat)
	logger.Println("历史聊天消息长度： ", len(msg.History))
	logger.Println(msg)
	return nil
}

// ResHeartbeat 心跳消息
func (cs *clientService) ResHeartbeat(s *session.Session, data proto.Message) error {
	//msg := data.(*message.ResHeartbeat)
	//logger.Println(msg)
	return nil
}
