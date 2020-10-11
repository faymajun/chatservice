package net

const (
	TCP_INIT = iota //待初始化
	TCP_AVAI        //可用
	TCP_STOP        //停止
)

const (
	MsgHeadSize = 4 //消息长度
	MsgIdSize   = 2 //消息Id
)

const (
	WriteQueLen = 64  // 写队列默认长度
	ReadQueLen  = 64  // 收队列默认长度
	bufferSize  = 512 // TCP接收缓冲区大小
)

const (
	DialTimeout      = 5 // 连接超时
	DefaultHeartBeat = 5 // 读超时
)

const (
	ClientWriteQue = 64 // Client写队列默认长度
	ClientReadQue  = 64 // Client收队列默认长度
)
