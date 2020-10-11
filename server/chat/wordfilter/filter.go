package wordfilter

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"regexp"
)

var logger = logrus.WithField("component", "filter")

// Filter
type Filter struct {
	trie  *Trie
	noise *regexp.Regexp
}

// NewFilter 创建一个Filter
func NewFilter(noisePattern string) *Filter {
	return &Filter{
		trie:  NewTrie(),
		noise: regexp.MustCompile(noisePattern),
	}
}

// LoadNetWordFile 加载网络Profanity word
func (f *Filter) LoadNetWordFile(url string) error {
	logger.Infof("加载网络Profanity word, %v", url)

	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	return f.load(rsp.Body)
}

// LoadLocalWordFile 加载本地 Profanity word 文件
func (f *Filter) LoadLocalWordFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return f.load(file)
}

// load 加载需要过滤的Profanity word
func (f *Filter) load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		f.trie.Add(string(line))
		logger.Infof("Profanity word：%s", string(line))
	}

	f.trie.BuildFailurePointer()
	return nil
}

// Replace 使用rep替换text中包含的Profanity word
func (f *Filter) Replace(text string, rep rune) string {
	return f.trie.Replace(text, rep)
}

// RemoveNoise 去除噪音
func (f *Filter) RemoveNoise(text string) string {
	return f.noise.ReplaceAllString(text, "")
}
