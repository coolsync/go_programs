package main

import (
	"fmt"
	"sync"
)

var (
	allData = make(map[string]string)
	rwm     sync.RWMutex
)

func get(key string) string {
	rwm.RLock()
	defer rwm.RUnlock()
	return allData[key]
}

func set(key, value string) {
	rwm.Lock()
	defer rwm.Unlock()
	allData[key] = value
}

func main() {
	set("a", "some data 1")
	set("b", "some data 2")

	res := get("a")
	// go get("b")
	// go get("a")

	fmt.Println(res)
	// time.Sleep(3*time.Second)
}
