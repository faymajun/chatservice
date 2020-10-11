package network

import (
	NET "cc-be-chat-test/go_modules/net"
	"context"
	"fmt"
	"net"
	"sync/atomic"
)

// TCPServer tcp服务器
type TCPServer struct {
	sid          uint32       // 唯一标识userSession
	listener     net.Listener // 监听器
	state        int32        // 状态
	addr         string       // 监听地址
	heartbeat    int64        // 心跳超时时间
	codecFactory func() NET.Codec
}

func (server *TCPServer) Close() error {
	if !atomic.CompareAndSwapInt32(&server.state, NET.TCP_AVAI, NET.TCP_STOP) {
		return NET.ErrDupClosed
	}
	return server.listener.Close()
}

func StartTcpServer(addr string, codecFactory func() NET.Codec, heartbeat int64, maxRecvSize uint32, maxRecvNum uint16) (*TCPServer, error) {
	return StartTcpProcessServer(addr, codecFactory, heartbeat, maxRecvSize, maxRecvNum, NET.WriteQueLen, NET.ReadQueLen)
}

func StartTcpProcessServer(addr string, codecFactory func() NET.Codec, heartbeat int64, maxRecvSize uint32, maxRecvNum uint16, wQueLen, rQueLen int) (*TCPServer, error) {
	lc := net.ListenConfig{KeepAlive: -1}
	listen, err := lc.Listen(context.Background(), "tcp4", addr)
	if err != nil {
		return nil, fmt.Errorf("tcp listen on %s failed, errstr:%s", addr, err.Error())
	}
	ts := &TCPServer{
		sid:          atomic.AddUint32(&serverId, 1),
		listener:     listen,
		state:        NET.TCP_AVAI,
		heartbeat:    heartbeat,
		codecFactory: codecFactory,
	}

	// check availability
	if wQueLen <= 0 {
		wQueLen = NET.WriteQueLen
	}
	if rQueLen <= 0 {
		rQueLen = NET.ReadQueLen
	}

	go ts.start(maxRecvSize, maxRecvNum, wQueLen, rQueLen)
	return ts, nil
}

func (server *TCPServer) start(maxRecvSize uint32, maxRecvNum uint16, wQueLen, rQueLen int) {
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			logger.Warnf("tcp accept failed:%s", err.Error())
			break
		}

		NET.New(conn, server.codecFactory(), maxRecvSize, maxRecvNum, wQueLen, rQueLen, server.heartbeat)
	}
	server.Close()
}
