package chat

import (
	"cc-be-chat-test/go_modules/message"
	"cc-be-chat-test/go_modules/net"
	"cc-be-chat-test/go_modules/net/session"
	"cc-be-chat-test/go_modules/server/chat/popular"
	"cc-be-chat-test/go_modules/server/chat/queue"
	"cc-be-chat-test/go_modules/server/chat/wordfilter"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

var logger = logrus.WithField("component", "room")

// 全局房间变量，因为时间关系，没有实现房间管理
var Room = newRoom(HistorySize)

// room 聊天房间
type room struct {
	history *queue.CircularQueue // 玩家聊天记录
	filter  *wordfilter.Filter   // 过滤Profanity word
	popular *popular.Popular     // 获取popular
}

// newRoom 返回一个聊天房间
func newRoom(historySize int) *room {
	return &room{
		history: queue.NewCircularQueue(historySize),
		filter:  wordfilter.NewFilter(NoisePattern),
		popular: popular.New(PopularTime),
	}
}

// Init 初始化
func (r room) Init() {
	// 获取网络文件
	//err := r.filter.LoadNetWordFile(ProfanityWordFile)
	// 获取本地文件
	err := r.filter.LoadLocalWordFile("go_modules/server/chat/profanityworddict/list.txt")
	if err != nil {
		panic(err)
	}
	logger.Infof("profanity word 加载完毕")
	go r.popular.Ticker()
	logger.Infof("聊天房间, Init成功")
}

// Chat 聊天广播
func (r room) Chat(s *session.Session, msg *message.Chat) error {
	if s == nil || msg == nil {
		return fmt.Errorf("Chat null error")
	}
	logger.Infof("收到聊天消息：%v", msg)

	// 处理 commands
	if strings.HasPrefix(msg.Content, Popular) {
		return r.wordOfPopular(s)
	} else if strings.HasPrefix(msg.Content, Stats) {
		return r.stats(s, msg.Content)
	}

	// 填充名字
	msg.Name = s.String(Name)

	// 过滤Profanity Word
	msg.Content = r.filter.Replace(r.filter.RemoveNoise(msg.Content), '*')

	// 保存词汇
	words := strings.Split(msg.Content, " ")
	r.popular.AddWord(words...)

	// 保存历史聊天数据
	r.history.EnQueue(msg)

	// 广播
	sessions := net.GetSessions()
	sessions.Range(func(_, v interface{}) bool {
		s := v.(*net.TCPSession)
		s.Send(message.MSGID_ResChatE, msg)
		return true
	})
	return nil
}

// Login client 进入聊天房间
func (r room) Login(s *session.Session) error {
	return s.Send(message.MSGID_ResHistoryChatE, &message.HistroyChat{History: r.history.GetAll()})
}
