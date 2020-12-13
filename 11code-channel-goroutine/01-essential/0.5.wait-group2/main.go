package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go sleep(&wg, time.Second*1)
	go sleep(&wg, time.Second*2)

	wg.Wait()

	fmt.Println("all go程 finished")
}

func sleep(wg *sync.WaitGroup, t time.Duration) {
	defer wg.Done()
	time.Sleep(t)
	fmt.Println("sleep finished")
}
