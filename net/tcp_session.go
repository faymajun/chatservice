package net

import (
	"bufio"
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net/coroutine"
	"cc-be-chat-test/go_modules/net/packet"
	"cc-be-chat-test/go_modules/net/session"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var sessionId uint32
var ErrDupClosed = errors.New("duplication close")
var sessions sync.Map

// TCPSession
type TCPSession struct {
	conn     net.Conn
	chWrite  chan *packet.SendMessage // 写通道
	ioReader *bufio.Reader            // 读缓冲
	state    int32                    // 状态
	sid      uint32                   // 唯一标识
	coder    Codec                    // 消息编解码

	UserSession *session.Session     // 用户session
	maxRecvSize uint32               //单个包最大收包大小(<=0无限制)
	maxRecvNum  uint16               //每秒最大收包数量(<=0无限制)
	curRecvTime int64                //收包时间
	curRecvNum  uint16               //当前收包数量
	routine     *coroutine.Coroutine //process 协程
}

func New(conn net.Conn, coder Codec, maxRecvSize uint32, maxRecvNum uint16, wQueLen, rQueLen int, heartBeat int64) *TCPSession {
	tcpSession := &TCPSession{
		conn:     conn,
		chWrite:  make(chan *packet.SendMessage, wQueLen),
		ioReader: bufio.NewReader(conn),
		state:    TCP_AVAI,
		sid:      atomic.AddUint32(&sessionId, 1),
		coder:    coder,

		maxRecvSize: maxRecvSize,
		maxRecvNum:  maxRecvNum,
		routine:     nil,
	}
	tcpSession.UserSession = session.New(tcpSession)
	sessions.Store(tcpSession.sid, tcpSession)
	routine := coroutine.New(rQueLen, int64(tcpSession.sid))
	tcpSession.SetRoutine(routine)

	if heartBeat == 0 {
		heartBeat = DefaultHeartBeat
	}
	go tcpSession.read(time.Duration(heartBeat) * time.Second)
	go tcpSession.write()

	return tcpSession
}

// SetRoutine 设置处理协程
func (s *TCPSession) SetRoutine(routine *coroutine.Coroutine) {
	s.routine = routine
}

// RemoteAddr
func (s *TCPSession) RemoteAddr() net.Addr { return s.conn.RemoteAddr() }

// Close
func (s *TCPSession) Close() error {
	if !atomic.CompareAndSwapInt32(&s.state, TCP_AVAI, TCP_STOP) {
		return ErrDupClosed
	}

	sessions.Delete(s.sid)
	close(s.chWrite)

	if s.routine != nil {
		s.routine.Close()
		s.routine = nil
	}

	return s.conn.Close()
}

func (s *TCPSession) isConnected() bool {
	return atomic.LoadInt32(&s.state) == TCP_AVAI
}

func (s *TCPSession) write() {
loop:
	for {
		select {
		case msg, ok := <-s.chWrite:
			if !ok {
				logrus.Debugf("session:%d write is closed", s.sid)
				break loop
			}
			payload, err := msg.Serialize()
			if err != nil {
				logrus.Errorf("session:%d serialize msg:%d failed:%s", s.sid, msg.MsgID, err)
				break loop
			}
			data := s.coder.Encode(msg.MsgID, payload)

			_, errWrite := s.conn.Write(data)
			if errWrite != nil {
				logrus.Errorf("session:%d  write err:%v", s.sid, errWrite)
				break loop
			}
		}
	}
	s.Close()
}

func (s *TCPSession) read(heartbeat time.Duration) {
	var head [MsgHeadSize]byte
	content := make([]byte, bufferSize)

	for {
		s.conn.SetReadDeadline(time.Now().Add(time.Second * heartbeat))
		if _, err := io.ReadFull(s.ioReader, head[:]); err != nil {
			if err != io.EOF {
				logrus.Warnf("session:%d read head is not io.EOF err:%s", s.sid, err)
			}

			logrus.Debugf("session:%d read is closed. err=%v", s.sid, err)
			break
		}

		size := binary.BigEndian.Uint32(head[:])
		if size > s.maxRecvSize && s.maxRecvSize > 0 {
			logrus.Warnf("session:%d size too big:%d, addr=%s", s.sid, size, s.RemoteAddr())
			break
		}
		if s.curRecvNum > s.maxRecvNum && s.maxRecvNum > 0 {
			logrus.Warnf("session:%d recv too many msg:%d, addr=%s", s.sid, s.curRecvNum, s.RemoteAddr())
			break
		}

		if size < MsgIdSize {
			logrus.Errorf("session:%d size too small:%d, addr=%s", s.sid, size, s.RemoteAddr())
			break
		}

		if size > uint32(len(content)) {
			content = make([]byte, size)
		}

		if _, err := io.ReadFull(s.ioReader, content[:size]); err != nil {
			if err != io.EOF {
				logrus.Warnf("session:%d read data err:%s, addr=%s", s.sid, err, s.RemoteAddr())
			}
			break
		}

		pack, err := s.coder.Decode(content[:size])
		if err != nil {
			logrus.Warnf("session:%d decode data err:%s, addr=%s", s.sid, err, s.RemoteAddr())
			break
		}

		if pack == packet.EmptyRecvPack {
			continue
		}

		if s.curRecvTime != time.Now().Unix() {
			s.curRecvTime = time.Now().Unix()
			s.curRecvNum = 0
		}
		s.curRecvNum++
		pack.Session = s.UserSession

		// 心跳直接回复
		if msg, ok := pack.Payload.(*message.ReqHeartbeat); ok {
			response := &message.ResHeartbeat{Uid: msg.Uid, ServerUnixTime: time.Now().Unix() * 1000}
			s.Send(message.MSGID_ResHeartbeatE, response)
			continue
		}

		if s.routine != nil {
			// 自身逻辑线程处理
			s.routine.PushPacket(&pack)
		}
	}
	s.Close()
}

func (s *TCPSession) send(msg *packet.SendMessage) error {
	if !s.isConnected() {
		return fmt.Errorf("session is closed: %s", s.RemoteAddr().String())
	}

	if len(s.chWrite) >= cap(s.chWrite) {
		logrus.Warnf("send buffer exceed, session close", s.RemoteAddr())

		s.Close()
		return fmt.Errorf("send buffer excced: %s", s.RemoteAddr())
	}

	s.chWrite <- msg
	return nil
}

// Send 发送数据
func (s *TCPSession) Send(msgid message.MSGID, pbmsg proto.Message) error {
	return s.send(&packet.SendMessage{MsgID: msgid, Payload: pbmsg})
}

// GetSessions 返回所有的Session
func GetSessions() sync.Map {
	return sessions
}
