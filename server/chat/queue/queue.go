package queue

import "cc-be-chat-test/go_modules/message"

// CircularQueue 循环队列：用于保存聊天最近数据
type CircularQueue struct {
	contents []*message.Chat
	index    int
	size     int
}

// NewCircularQueue 新建一个循环队列
func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		contents: make([]*message.Chat, size),
		index:    0,
		size:     size,
	}
}

// EnQueue 聊天数据入队
func (q *CircularQueue) EnQueue(chat *message.Chat) {
	q.contents[q.index] = chat
	q.index = (q.index + 1) % q.size
}

// GetAll 获取队列所有
func (q *CircularQueue) GetAll() []*message.Chat {
	// 不足
	if q.contents[q.index] == nil {
		return q.contents[0:q.index]
	} else {
		// 按照顺序填充数据
		var res = make([]*message.Chat, 0)
		res = append(res, q.contents[q.index:q.size]...)
		res = append(res, q.contents[0:q.index]...)
		return res
	}
}
