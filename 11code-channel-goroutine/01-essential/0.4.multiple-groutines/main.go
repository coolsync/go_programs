package main

import (
	"fmt"
	"time"
)

func excute(id int) {
	fmt.Printf("id: %d\n", id)
}

func main() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		go excute(i)
	}
	time.Sleep(1*time.Second)
	fmt.Println("end")
}