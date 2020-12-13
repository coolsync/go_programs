package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go send(ch)
	go receive(ch)

	go send2(ch)
	fmt.Printf("main go receive send go val: %d\n", <-ch)
	time.Sleep(time.Second * 2)
}

func send2(ch chan int) {

	ch <- 12
}

func send(ch chan int) {
	time.Sleep(time.Second * 1)
	fmt.Println("解除send blocking")

	ch <- 12
}

func receive(ch chan int) {
	// fmt.Println("解除blocking")

	n := <-ch
	fmt.Printf("接受send go 的值 %d\n", n)
}

// func send(ch chan int) {
// 	ch <- 12
// 	fmt.Println("发送完毕")
// }

// func receive(ch chan int) {
// 	time.Sleep(time.Second * 1)
// 	fmt.Println("解除blocking")

// 	_ = <-ch
// 	return
// }

