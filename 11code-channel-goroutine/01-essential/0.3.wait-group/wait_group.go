package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		// Int returns a non-negative pseudo-random int from the default Source.
		r1 := rand.Int()
		// Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source.
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)

	}
}

func f1(i int) {
	defer wg.Done()
	// 播种随机数种子
	// rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond *time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}

	wg.Wait()
}
