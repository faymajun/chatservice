package net

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net/route"
	"encoding/binary"
	"fmt"

	"cc-be-chat-test/go_modules/net/packet"
	"github.com/golang/protobuf/proto"
)

var OrdinaryCodecFactory = func() Codec { return NewOrdinaryCoder() }

// Codec
type Codec interface {
	Encode(msgid message.MSGID, payload []byte) []byte
	Decode(data []byte) (packet.RecvMessage, error)
}

// OrdinaryCoder 通过解码器
type OrdinaryCoder struct {
}

// NewOrdinaryCoder
func NewOrdinaryCoder() *OrdinaryCoder {
	return &OrdinaryCoder{}
}

// Encode
func (OrdinaryCoder) Encode(msgid message.MSGID, payload []byte) []byte {
	data := make([]byte, MsgHeadSize+MsgIdSize+len(payload))
	binary.BigEndian.PutUint32(data, uint32(MsgIdSize+len(payload)))
	binary.BigEndian.PutUint16(data[MsgHeadSize:], uint16(msgid))
	copy(data[MsgHeadSize+MsgIdSize:], payload)
	return data
}

// Decode
func (o OrdinaryCoder) Decode(data []byte) (packet.RecvMessage, error) {
	msgid := binary.BigEndian.Uint16(data[0:MsgIdSize])
	handler, err := route.FindHandler(message.MSGID(msgid))
	if err != nil {
		return packet.EmptyRecvPack, err
	}

	size := len(data)

	pbMsg := handler.Instance()
	pbbuf := proto.NewBuffer(data[MsgIdSize:])
	if err := pbbuf.Unmarshal(pbMsg); err != nil {
		return packet.EmptyRecvPack, fmt.Errorf("解码错误: MsgId=%d, Length=%v Error=%v", msgid, size-MsgIdSize, err)
	}

	return packet.RecvMessage{Handler: handler, Payload: pbMsg}, nil
}
