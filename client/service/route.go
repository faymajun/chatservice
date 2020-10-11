package service

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net/route"
)

// 路由表
var RouteTable = []route.Stub{
	{
		MsgID:   message.MSGID_ResLoginE,
		Handler: ClientService.Login,
		Payload: new(message.ReqLogin),
	},
	{
		MsgID:   message.MSGID_ResChatE,
		Handler: ClientService.Chat,
		Payload: new(message.Chat),
	},
	{
		MsgID:   message.MSGID_ResHistoryChatE,
		Handler: ClientService.History,
		Payload: new(message.HistroyChat),
	},
	{
		MsgID:   message.MSGID_ResHeartbeatE,
		Handler: ClientService.ResHeartbeat,
		Payload: new(message.ResHeartbeat),
	},
}
