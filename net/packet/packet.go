package packet

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net/route"
	"cc-be-chat-test/go_modules/net/session"
	"errors"

	"github.com/golang/protobuf/proto"
)

var (
	ErrUnsupportedPayload = errors.New("发送数据包不支持的Payload类型")
)

type (
	// SendMessage 发送消息结构体
	SendMessage struct {
		MsgID   message.MSGID // 消息ID
		Payload interface{}   // 消息内容
		// 消息可以是proto.Message或[]byte, 如
		// 果是[]byte表示广播消息, 已经提前序列化好
	}

	// RecvMessage 接受消息结构体
	RecvMessage struct {
		Session *session.Session    // session
		Handler *route.LogicHandler // 处理函数
		Payload proto.Message       // 消息内容
	}
)

var EmptyRecvPack = RecvMessage{}

// Serialize 返回序列化后的数据
func (p *SendMessage) Serialize() ([]byte, error) {
	switch v := p.Payload.(type) {
	case []byte:
		return v, nil
	case proto.Message:
		return proto.Marshal(v)
	default:
		return nil, ErrUnsupportedPayload
	}
}
