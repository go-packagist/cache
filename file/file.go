package file

import (
	"github.com/go-packagist/cache"
	"time"
)

type Store struct {
	options *Options
}

type Options struct {
	path string
}

func New(options *Options) cache.Cache {
	store := &Store{
		options: options,
	}

	store.GC()

	return store
}

func (s *Store) Put(key string, value interface{}, expire time.Duration) error {
	// todo

	return nil
}

func (s *Store) Get(key string) (interface{}, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) Has(key string) bool {
	// TODO implement me
	panic("implement me")
}

func (s *Store) Remember(key string, fc func() interface{}, expire time.Duration) (interface{}, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) GC() error {
	// TODO

	return nil
}
