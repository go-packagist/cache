package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemorySetOrGet(t *testing.T) {
	memory := NewMemory()

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

func TestMemoryMuti(t *testing.T) {
	memory, memory2 := NewMemory(), NewMemory()

	memory.Put("a", "aaa", time.Second*1)
	assert.Equal(t, "aaa", memory.Get("a").Value())
	assert.Equal(t, nil, memory2.Get("b").Value())

	memory2.Put("b", "bbb", time.Second*1)
	assert.Equal(t, "bbb", memory2.Get("b").Value())
}
