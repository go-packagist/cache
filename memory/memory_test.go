package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var memory = New()

func TestMemoryBase(t *testing.T) {
	memory.Put("a", "aaa", time.Second*1)
	a, _ := memory.Get("a")

	time.Sleep(time.Second * 2)
	aa, _ := memory.Get("a")

	b, _ := memory.Remember("b", func() interface{} {
		return "bbb"
	}, time.Second*1)

	assert.Equal(t, "aaa", a)
	assert.Equal(t, true, memory.Has("a"))
	assert.Equal(t, false, memory.Has("b"))
	assert.Equal(t, nil, aa)
	assert.Equal(t, "bbb", b)
}
