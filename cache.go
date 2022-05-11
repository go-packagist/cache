package cache

import "time"

type Cache interface {
	Get(key string) (interface{}, error)
	Put(key string, value interface{}, expire time.Duration) error
	Has(key string) bool
	Remember(key string, fc func() interface{}, expire time.Duration) (interface{}, error)
	GC() error
}

var stores = make(map[string]Cache)

func Configure(name string, store Cache) {
	stores[name] = store
}

func Store(name string) Cache {
	return stores[name]
}
