package network

import (
	"cc-be-chat-test/go_modules/net"
	"github.com/sirupsen/logrus"
)

var logger = logrus.WithField("module", "network")

func StopTcpSession() {
	logger.Println("<<<sessions is stop start>>>")
	var sessions = net.GetSessions()
	sessions.Range(func(k, v interface{}) bool {
		if s, ok := v.(*net.TCPSession); ok {
			s.Close()
		}
		return true
	})
	logger.Println("<<<sessions is stop over>>>")
}

var serverId uint32
