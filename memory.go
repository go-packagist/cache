package cache

import (
	"errors"
	"time"
)

type memoryStore struct {
	data map[string]*memoryData
}

type memoryData struct {
	key    string
	value  interface{}
	expire time.Time
}

func NewMemory() Cache {
	store := &memoryStore{
		data: make(map[string]*memoryData),
	}

	store.GC()

	return store
}

func (s *memoryStore) Put(key string, value interface{}, expire time.Duration) error {
	s.data[key] = &memoryData{
		key:    key,
		value:  value,
		expire: time.Now().Add(expire),
	}

	return nil
}

func (s *memoryStore) Get(key string) *Result {
	data, ok := s.data[key]

	if ok {
		return &Result{data.value, nil}
	}

	return &Result{nil, errors.New("key not found")}
}

func (s *memoryStore) Has(key string) bool {
	_, ok := s.data[key]

	return ok
}

func (s *memoryStore) Remember(key string, fc func() interface{}, expire time.Duration) *Result {
	if !s.Has(key) {
		s.Put(key, fc(), expire)
	}

	return s.Get(key)
}

func (s *memoryStore) GC() error {
	go func() {
		for {
			for key, data := range s.data {
				if data.expire.Before(time.Now()) {
					delete(s.data, key)
				}
			}
		}
	}()

	return nil
}
