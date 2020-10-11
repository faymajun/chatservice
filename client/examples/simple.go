package main

import (
	"cc-be-chat-test/go_modules/client/network"
	"cc-be-chat-test/go_modules/client/service"
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net"
	"cc-be-chat-test/go_modules/net/route"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Popular    = "/popular"
	Stats      = "/stats ann"
	userName   = "bob"
	//Stats      = "/stats bob"  切换name stats 可方便测试
	//userName   = "ann"
	serverPort = ":660"
	heartbeat  = 5 // 心跳间隔
)

func main() {
	route.RegisterTable(service.RouteTable) // 注册消息处理函数
	var serverAddr = serverPort             // 服务器地址

	server, err := network.StartTcpClient(serverAddr, net.OrdinaryCodecFactory)
	if err != nil {
		panic(err)
	}

	go func() {
		// 登录
		server.Send(message.MSGID_ReqLoginE, &message.ReqLogin{Name: userName})
		for {
			// 聊天
			server.Send(message.MSGID_ReqChatE, &message.Chat{Content: "hellboy 222 "})
			server.Send(message.MSGID_ReqChatE, &message.Chat{Content: "hello bitch world! 333"})
			server.Send(message.MSGID_ReqChatE, &message.Chat{Content: "hello bi@#tch world! 444"}) //噪音词
			time.Sleep(5 * time.Second)
			// 命令
			server.Send(message.MSGID_ReqChatE, &message.Chat{Content: Popular})
			server.Send(message.MSGID_ReqChatE, &message.Chat{Content: Stats})
		}
	}()

	logrus.Warnf("Chat Client启动成功, 服务器地址=%s, 心跳间隔=%d秒 ,%d", serverAddr, heartbeat, time.Now().UnixNano())

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
	// 断开与服务器的连接
	server.Close()
	logrus.Infof("Client shutdown Finish!!!")
}
