package service

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net/route"
)

// 前后端服务器都需要监听的路由表
var RouteTable = []route.Stub{
	{
		MsgID:   message.MSGID_ReqLoginE,
		Handler: ClientService.Login,
		Payload: new(message.ReqLogin),
	},
	{
		MsgID:   message.MSGID_ReqChatE,
		Handler: ClientService.Chat,
		Payload: new(message.Chat),
	},
	{
		MsgID:   message.MSGID_ReqHeartbeatE,
		Handler: ClientService.Heartbeat,
		Payload: new(message.ReqHeartbeat),
	},
}
