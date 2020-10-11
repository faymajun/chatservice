package route

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net/session"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"reflect"
	"runtime/debug"
)

type (
	// Stub 路由表
	Stub struct {
		MsgID   message.MSGID
		Handler HandlerFunc
		Payload proto.Message
	}

	// HandlerFunc 消息处理函数
	HandlerFunc func(*session.Session, proto.Message) error

	// LogicHandler 消息处理器
	LogicHandler struct {
		msgid message.MSGID // 消息ID
		fn    HandlerFunc   // 回调函数
		typ   reflect.Type  // 消息类型
	}
)

// Handle 使用反序列化好的消息和client对应的session调用消息处理函数
func (h *LogicHandler) Handle(s *session.Session, payload proto.Message) error {
	// 防止逻辑处理过程中panic导致逻辑线程crash
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("Handler panic: MsgID=%s, Error=%v", h.msgid, err)
			fmt.Fprintln(logrus.StandardLogger().Out, string(debug.Stack()))
		}
	}()

	// logic processor
	return h.fn(s, payload)
}

// Instance 生成一个新的消息实例
func (h *LogicHandler) Instance() proto.Message {
	return reflect.New(h.typ).Interface().(proto.Message)
}

// MsgID
func (h *LogicHandler) MsgID() message.MSGID {
	return h.msgid
}
