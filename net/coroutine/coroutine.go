package coroutine

import (
	"cc-be-chat-test/go_modules/net/packet"
	"errors"
	"fmt"
	"sync/atomic"
)

// Coroutine 处理接受到的网络消息
type Coroutine struct {
	chRead chan *packet.RecvMessage // 读通道
	close  int32                    // 状态
	userID int64                    // 用户id
}

// New 创建一个新的协程
func New(queCap int, userId int64) *Coroutine {
	c := &Coroutine{
		chRead: make(chan *packet.RecvMessage, queCap),
		close:  0,
		userID: userId,
	}

	go c.Process()
	return c
}

func (c *Coroutine) isClose() bool {
	return atomic.LoadInt32(&c.close) == 1
}

// Close
func (c *Coroutine) Close() error {
	if !atomic.CompareAndSwapInt32(&c.close, 0, 1) {
		return errors.New("duplication close")
	}

	close(c.chRead)
	return nil
}

// Process 消息处理
func (c *Coroutine) Process() {
loop:
	for {
		select {
		case pack, ok := <-c.chRead:
			if !ok {
				c.chRead = nil
				break loop
			}
			process(pack)
		}
	}

	c.Close()
}

// 添加接受到的消息
func (c *Coroutine) PushPacket(packet *packet.RecvMessage) error {
	if c.isClose() {
		return nil
	}

	c.chRead <- packet
	return nil
}

func process(pack *packet.RecvMessage) error {
	var (
		session = pack.Session
		handler = pack.Handler
		pbMsg   = pack.Payload
	)

	if err := handler.Handle(session, pbMsg); err != nil {
		return fmt.Errorf(fmt.Sprintf("Handle %v error: %v", handler.MsgID(), err))
	}
	return nil
}
