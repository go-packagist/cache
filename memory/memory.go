package memory

import (
	"errors"
	"github.com/go-packagist/cache"
	"time"
)

type Store struct {
	data map[string]*Data
}

type Data struct {
	key    string
	value  interface{}
	expire time.Time
}

func New() cache.Cache {
	store := &Store{
		data: make(map[string]*Data),
	}

	store.GC()

	return store
}

func (s *Store) Put(key string, value interface{}, expire time.Duration) error {
	s.data[key] = &Data{
		key:    key,
		value:  value,
		expire: time.Now().Add(expire),
	}

	return nil
}

func (s *Store) Get(key string) (interface{}, error) {
	data, ok := s.data[key]

	if ok {
		return data.value, nil
	}

	return nil, errors.New("key not found")
}

func (s *Store) Has(key string) bool {
	_, ok := s.data[key]

	return ok
}

func (s *Store) Remember(key string, fc func() interface{}, expire time.Duration) (interface{}, error) {
	if !s.Has(key) {
		s.Put(key, fc(), expire)
	}

	return s.Get(key)
}

func (s *Store) GC() error {
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
