package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

// This is a more expressive name to the variables related to the cache
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func LookupV2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	fmt.Println(Lookup("item"))
}
