package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var memory = NewMemory()

func TestMemorySetOrGet(t *testing.T) {
	// Put
	memory.Put("a", "aaa", time.Second*1)

	// Get
	result := memory.Get("a")
	assert.Equal(t, "aaa", result.Value())
	assert.Equal(t, "aaa", result.Val)
	assert.False(t, result.IsError())
	assert.Equal(t, nil, result.Error())

	// Expire(GC)
	time.Sleep(time.Second * 2)
	assert.Equal(t, nil, memory.Get("a").Value())
	assert.True(t, memory.Get("a").IsError())
}
