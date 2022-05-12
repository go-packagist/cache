package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var memory = NewMemory()

func TestMemory(t *testing.T) {
	memory.Put("a", "aaa", time.Second*1)
	a, _ := memory.Get("a")
	assert.Equal(t, "aaa", a)
	assert.Equal(t, true, memory.Has("a"))
	assert.Equal(t, false, memory.Has("b"))

	time.Sleep(time.Second * 2)
	aa, _ := memory.Get("a")
	assert.Equal(t, nil, aa)

	b, _ := memory.Remember("b", func() interface{} {
		return "bbb"
	}, time.Second*1)
	assert.Equal(t, true, memory.Has("b"))
	assert.Equal(t, "bbb", b)
}
