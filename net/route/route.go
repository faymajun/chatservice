package route

import (
	"cc-be-chat-test/go_modules/message"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// handlers 消息ID对应的回调处理函数
var handlers = [message.MSGID_MAX_COUNT]LogicHandler{}

// register 注册消息回调函数
func register(msgid message.MSGID, handler HandlerFunc, payload proto.Message, isRaw ...bool) {
	if int(msgid) >= len(handlers) {
		log.Errorf("消息ID超出范围: 最大值=%d, 当前值=%d, ID=%s", len(handlers), int(msgid), msgid.String())
		return
	}

	if handler == nil {
		log.Error("消息处理函数不可为空")
		return
	}

	// 检查是都已经注册
	if h := handlers[msgid]; h.fn != nil {
		log.Errorf("消息ID已经注册: %s", msgid.String())
		return
	}

	payloadType := reflect.TypeOf(payload)
	if payloadType.Kind() != reflect.Ptr {
		log.Errorf("消息类型必须为一个指针: %s", msgid.String())
		return
	}
	handlers[msgid] = LogicHandler{msgid: msgid, fn: handler, typ: payloadType.Elem()}
	log.Infof("注册消息: MsgID=%s", msgid)
}

// RegisterTable 注册消息处理函数
func RegisterTable(table []Stub) {
	for _, item := range table {
		register(item.MsgID, item.Handler, item.Payload)
	}
}

// FindHandler 获取消息ID对应的处理器函数
func FindHandler(msgid message.MSGID) (*LogicHandler, error) {
	if int(msgid) >= len(handlers) {
		return nil, errors.Errorf("message id over range: %s", msgid.String())
	}

	if h := handlers[msgid]; h.fn == nil {
		return nil, errors.Errorf("handler not found: %s", msgid.String())
	}

	return &handlers[msgid], nil
}
