package main

import (
	"fmt"
	"sync"
)



/* 
var (
	mu sync.Mutex
	m  = make(map[string]string)
)

func lookup(key string) string {
	mu.Lock()
	defer mu.Unlock()
	v := m[key]

	return v
} 
*/

// 下面这个版本在功能上是一致的，两个包级别的变量放在了cache这个struct一组内：
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookup(key string) string {
	cache.Lock()
	defer cache.Unlock()

	v := cache.mapping[key]

	return v
}

func main() {
	cache.mapping = map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	}

	for k := range cache.mapping {
		fmt.Println(lookup(k))
	}
}
