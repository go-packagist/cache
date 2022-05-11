package redis

import (
	"errors"
	"github.com/go-packagist/cache"
	"github.com/go-redis/redis/v8"
	"time"
)

type Store struct {
	client *redis.Client
}

func New(client *redis.Client) cache.Cache {
	store := &Store{
		client: client,
	}

	store.GC()

	return store
}

func (s *Store) Put(key string, value interface{}, expire time.Duration) error {
	// TODO
	return nil
}

func (s *Store) Get(key string) (interface{}, error) {
	// TODO
	return nil, errors.New("key not found")
}

func (s *Store) Has(key string) bool {
	// TODO
	return true
}

func (s *Store) Remember(key string, fc func() interface{}, expire time.Duration) (interface{}, error) {
	if !s.Has(key) {
		s.Put(key, fc(), expire)
	}

	return s.Get(key)
}

func (s *Store) GC() error {
	return nil
}
