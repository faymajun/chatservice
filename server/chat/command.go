package chat

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net"
	"cc-be-chat-test/go_modules/net/session"
	"fmt"
	"strings"
	"time"
)

// wordOfPopular 发送clint popular word
func (r room) wordOfPopular(s *session.Session) error {
	return s.Send(message.MSGID_ResChatE, &message.Chat{Content: r.popular.Most()})
}

// stats 发送client 用户的stats
func (r room) stats(s *session.Session, content string) error {
	text := strings.Split(content, " ")
	if len(text) < 2 {
		return fmt.Errorf("stats context err: %s", content)
	}

	// todo 因为时间关系，没有建立 map[name]session的关系，从而优化时间到O(1)
	var loginTime int64 = 0
	var name = text[1]
	sessions := net.GetSessions()
	sessions.Range(func(_, v interface{}) bool {
		s := v.(*net.TCPSession)
		if s.UserSession.String(Name) == name {
			loginTime = s.UserSession.Int64(LoginTime)
		}
		return true
	})

	if loginTime == 0 {
		return s.Send(message.MSGID_ResChatE, NoneTimeStatsMsg)
	} else {
		lt := time.Unix(loginTime, 0)
		diffTime := time.Now().Sub(lt)

		stats := fmt.Sprintf("%02dd %02dh %02dm %02ds",
			int(diffTime.Hours())/24, int(diffTime.Hours())%24, int(diffTime.Minutes()), int(diffTime.Seconds()))
		return s.Send(message.MSGID_ResChatE, &message.Chat{Content: stats})
	}
}

var NoneTimeStatsMsg = &message.Chat{Content: "00d 00h 00m 00s"}
