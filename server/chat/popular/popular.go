package popular

import (
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

//// Word Popular使用计算结构体
//type Word struct {
//	context string  // 内容
//	count int  // 数量
//}

var logger = logrus.WithField("component", "Popular")

// Popular 获取最后5秒内，使用最后的word
// 每秒更新，lastData = lastData + cur - preFifth(前面第五的数据)
type Popular struct {
	preData      []map[string]int // 往前5秒，每秒的数据
	size         int              // 当前保存了多少秒的数据
	cap          int              // cap of preData 计算最后多少秒的数据
	cur          map[string]int   // 当前这秒内的数据
	lastData     map[string]int   // 前五秒的的数据总和
	theMost      string           // 当前最多的数据
	sync.RWMutex                  // mutex
}

// New 返回一个popular
func New(saveTime int) *Popular {
	p := &Popular{
		preData:  make([]map[string]int, 0),
		size:     0,
		cap:      saveTime,
		cur:      make(map[string]int),
		lastData: make(map[string]int),
	}

	return p
}

// Ticker 定时时钟
func (p *Popular) Ticker() {
	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				p.UpdateBySecond()
			}
		}
	}()
}

// UpdateBySecond 每秒更新
// 更新 lastFive， theMost,
func (p *Popular) UpdateBySecond() {
	//logger.Infof("Popular UpdateBySecond")
	p.Lock()
	defer p.Unlock()

	// todo 因为时间关系，p.lastData 没有使用循环队列来进行优化
	// 移除老数据
	if p.size >= p.cap {
		//logger.Infof("移除老数据：%v", p.preData[0])
		for word, num := range p.preData[0] {
			p.lastData[word] -= num
		}
		p.preData = p.preData[1:] // 移除头部最老的数据
		p.size--
	}

	// 增加最新数据
	for word, num := range p.cur {
		p.lastData[word] += num
	}

	p.preData = append(p.preData, p.cur) // 添加尾部新数据
	p.size++
	p.cur = make(map[string]int)

	// 计算most
	var mostWord = ""
	var mostNum = 0
	for word, num := range p.lastData {
		if num > mostNum {
			mostNum = num
			mostWord = word
		}
	}
	p.theMost = mostWord
}

// Most 返回当前最多的数据
func (p *Popular) Most() string {
	return p.theMost
}

// AddWord 添加word
func (p *Popular) AddWord(words ...string) {
	//logger.Infof("AddWord: %s 成功", words)
	p.Lock()
	defer p.Unlock()

	for _, word := range words {
		if strings.Contains(word, "*") {
			continue
		}
		//logger.Infof("添加word: %s 成功", word)
		p.cur[word]++
	}
}
