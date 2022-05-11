# cache

## Installation

```bash
go get github.com/go-packagist/cache
```

## Usage

```golang
package main

import (
	"cache/cache"
	"cache/cache/memory"
	"fmt"
	"time"
)

func main() {
	cache.Configure("memory", memory.New())
	cache.Configure("memory2", memory.New())

	cache.Store("memory").Put("a", "2", time.Second*1)
	cache.Store("memory2").Put("aa", "2", time.Second*1)

	fmt.Println(cache.Store("memory").Get("a"))
	fmt.Println(cache.Store("memory2").Get("aa"))
	fmt.Println(cache.Store("memory2").Get("a"))

	time.Sleep(time.Second * 2)

	fmt.Println(cache.Store("memory").Get("a"))
	fmt.Println(cache.Store("memory2").Get("aa"))

	fmt.Println("asdfasdf")
}
```