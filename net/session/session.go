package session

import (
	"cc-be-chat-test/go_modules/message"
	"errors"
	"github.com/golang/protobuf/proto"
	"net"
	"sync"
	"sync/atomic"
)

// NetworkEntity 底层网络实例
type NetworkEntity interface {
	Send(msgid message.MSGID, pbmsg proto.Message) error
	Close() error
	RemoteAddr() net.Addr
}

var (
	// ErrIllegalUserID 不合法userId
	ErrIllegalUserID = errors.New("illegal userID")
	sid int64
)

// Session 保存连接client与sever的连接session
// 当网络断开，所有数据将被释放。
type Session struct {
	id           int64                  // session 全局唯一 id
	userID       int64                  // 绑定 user ID
	entity       NetworkEntity          // 底层网络实例
	sync.RWMutex                        // mutex
	data         map[string]interface{} // data store
}

// New 返回一个 网络连接session
// NetworkEntity 是底层网络实例
func New(entity NetworkEntity) *Session {
	return &Session{
		id:     atomic.AddInt64(&sid, 1),
		entity: entity,
		data:   make(map[string]interface{}),
	}
}

// NetworkEntity  返回一个 NetworkEntity 底层网络实例
func (s *Session) NetworkEntity() NetworkEntity {
	return s.entity
}

// Send 发送数据
func (s *Session) Send(msgid message.MSGID, pbmsg proto.Message) error {
	return s.entity.Send(msgid, pbmsg)
}

// ID 返回session id
func (s *Session) ID() int64 {
	return s.id
}

// UserID 返回当前session绑定的用户ID
func (s *Session) UserID() int64 {
	return atomic.LoadInt64(&s.userID)
}

// Bind 绑定的用户ID
func (s *Session) Bind(userID int64) error {
	if userID < 1 {
		return ErrIllegalUserID
	}

	atomic.StoreInt64(&s.userID, userID)
	return nil
}

// Close 关闭当前会话
func (s *Session) Close() {
	s.entity.Close()
}

// RemoteAddr 返回 remote network address.
func (s *Session) RemoteAddr() net.Addr {
	return s.entity.RemoteAddr()
}

// Remove 移除数据
func (s *Session) Remove(key string) {
	s.Lock()
	defer s.Unlock()

	delete(s.data, key)
}

// Set 保存数据
func (s *Session) Set(key string, value interface{}) {
	s.Lock()
	defer s.Unlock()

	s.data[key] = value
}

// HasKey 是否存在key
func (s *Session) HasKey(key string) bool {
	s.RLock()
	defer s.RUnlock()

	_, has := s.data[key]
	return has
}

// Clear 清除所有数据
func (s *Session) Clear() {
	s.Lock()
	defer s.Unlock()

	s.userID = 0
	s.data = map[string]interface{}{}
}

// String 返回key指向的string数据
func (s *Session) String(key string) string {
	s.RLock()
	defer s.RUnlock()

	v, ok := s.data[key]
	if !ok {
		return ""
	}

	value, ok := v.(string)
	if !ok {
		return ""
	}
	return value
}

// Int64 返回int64
func (s *Session) Int64(key string) int64 {
	s.RLock()
	defer s.RUnlock()

	v, ok := s.data[key]
	if !ok {
		return 0
	}

	value, ok := v.(int64)
	if !ok {
		return 0
	}
	return value
}
