package main

import (
	"cc-be-chat-test/go_modules/net"
	"cc-be-chat-test/go_modules/net/route"
	"cc-be-chat-test/go_modules/server/chat"
	"cc-be-chat-test/go_modules/server/network"
	"cc-be-chat-test/go_modules/server/service"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	serverPort       = ":660" // 监听地址
	heartbeat  int64 = 5      // 心跳间隔
)

func main() {
	route.RegisterTable(service.RouteTable) // 注册消息处理函数

	// 聊天房间初始化
	chat.Room.Init()

	// 监听端口，设置心跳，设置包最大值，设置每秒最多包数量，开启每个客户端一个处理协程
	server, err := network.StartTcpServer(serverPort, net.OrdinaryCodecFactory, heartbeat, 1024, 20)
	if err != nil {
		panic(err)
	}

	logrus.Warnf("Chat服务器启动成功, TCP监听地址=%s, 心跳间隔=%d秒 ,%d", serverPort, heartbeat, time.Now().UnixNano())

	// 等待退出信号
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig, _ := <-stopChan:
		logrus.Infof("<<<==================>>>")
		logrus.Infof("<<<stop process by:%v>>>", sig)
		logrus.Infof("<<<==================>>>")
		break
	}

	signal.Stop(stopChan)
	close(stopChan)
	// 关闭服务器监听，阻止新连接
	server.Close()
	// 关闭客户端连接
	network.StopTcpSession()

	logrus.Infof("Server shutdown Finish!!!")
}
