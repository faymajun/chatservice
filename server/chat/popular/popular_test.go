package popular

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPopular(t *testing.T) {
	p := New(5)
	go p.Ticker()
	assert := assert.New(t)
	var tests = []struct {
		input    []string
		expected string
	}{
		{[]string{"a", "b", "c", "c"}, ""}, // 前面两次都一次添加
		{[]string{"a", "a"}, "a"},          // 前面两次都一次添加
		{[]string{"b", "b", "b", "b"}, "b"},
		{[]string{"c", "c", "c", "c"}, "c"},
		{[]string{"e"}, "c"},
		{[]string{"f"}, "c"},
		{[]string{"g"}, "b"},
		{[]string{""}, "c"},
	}
	for _, test := range tests {
		p.AddWord(test.input...)
		fmt.Println("add word after")
		time.Sleep(900 * time.Millisecond)
		fmt.Println("check..")
		assert.Equal(p.Most(), test.expected)
	}
}
