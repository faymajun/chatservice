package network

import (
	"cc-be-chat-test/go_modules/message"
	NET "cc-be-chat-test/go_modules/net"
	"cc-be-chat-test/go_modules/net/session"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
)

// TCPClient client tcp连接
type TCPClient struct {
	session      *NET.TCPSession  // Session
	state        int32            // 状态
	addr         string           // 对端地址
	codecFactory func() NET.Codec // 消息编解码
	heartClose   chan struct{}    // 心跳退出信号
}

// Close
func (client *TCPClient) Close() error {
	if !atomic.CompareAndSwapInt32(&client.state, NET.TCP_AVAI, NET.TCP_STOP) {
		return NET.ErrDupClosed
	}

	close(client.heartClose)
	if client.session == nil {
		return nil
	}
	return client.session.Close()
}

func (client *TCPClient) isRunning() bool {
	return atomic.LoadInt32(&client.state) == NET.TCP_AVAI
}

// StartTcpClient
func StartTcpClient(addr string, codecFactory func() NET.Codec) (*TCPClient, error) {
	client := &TCPClient{
		state:        NET.TCP_AVAI,
		addr:         addr,
		codecFactory: codecFactory,
		heartClose:   make(chan struct{}),
	}

	_, err := client.dial()
	return client, err
}

func (client *TCPClient) dial() (conn net.Conn, err error) {
	conn, err = net.DialTimeout("tcp4", client.addr, time.Second*NET.DialTimeout)
	if err != nil {
		return
	}

	tcpSession := NET.New(conn, client.codecFactory(), 0, 0, NET.ClientWriteQue, NET.ClientReadQue, 0)
	client.session = tcpSession
	go client.heartbeat()
	return
}

func (client *TCPClient) heartbeat() {
	t := time.NewTicker(1 * time.Second)
	p := &message.ReqHeartbeat{}
	defer t.Stop()
loop:
	for {
		select {
		case <-client.heartClose:
			break loop

		case <-t.C:
			if !client.isRunning() {
				break loop
			}
			client.Send(message.MSGID_ReqHeartbeatE, p)
		}
	}
	logrus.Infof("远程服务器心跳线程退出, Addr=%s", client.addr)
}

// Send 发送消息
func (client *TCPClient) Send(msgid message.MSGID, pbmsg proto.Message) error {
	if !client.isRunning() {
		return fmt.Errorf("client is not running")
	}
	sess := client.session
	if sess != nil {
		return sess.Send(msgid, pbmsg)
	}
	return fmt.Errorf("client session is nil")
}

// RemoteAddr
func (client *TCPClient) RemoteAddr() net.Addr {
	return client.session.RemoteAddr()
}

// Addr
func (client *TCPClient) Addr() string {
	return client.addr
}

// Session 返回session
func (client *TCPClient) Session() *session.Session {
	return client.session.UserSession
}
