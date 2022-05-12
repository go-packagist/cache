package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemoryStore_SetAndGet(t *testing.T) {
	memory := NewMemoryStore()

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

func TestMemoryStore_Muti(t *testing.T) {
	memory, memory2 := NewMemoryStore(), NewMemoryStore()

	memory.Put("aa", "aaa", time.Second*1)
	assert.Equal(t, "aaa", memory.Get("aa").Value())
	assert.Equal(t, nil, memory2.Get("bb").Value())

	memory2.Put("bb", "bbb", time.Second*1)
	assert.Equal(t, "bbb", memory2.Get("bb").Value())
}
