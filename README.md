# cache

## TODO

- [x] Memory
- [ ] Redis
- [ ] Memcached
- [ ] FileSystem

## Installation

```bash
go get github.com/go-packagist/cache
```

## Usage

```golang
package main

import (
	"fmt"
	"github.com/go-packagist/cache"
	"time"
)

func main() {
	cache.Configure("memory", cache.NewMemory())
	cache.Configure("memory2", cache.NewMemory())

	cache.Store("memory").Put("a", "2", time.Second*1)
	cache.Store("memory2").Put("aa", "2", time.Second*1)

	fmt.Println(cache.Store("memory").Get("a"))
	fmt.Println(cache.Store("memory2").Get("aa"))
	fmt.Println(cache.Store("memory2").Get("a"))

	time.Sleep(time.Second * 2)

	fmt.Println(cache.Store("memory").Get("a"))
	fmt.Println(cache.Store("memory2").Get("aa"))
}
```