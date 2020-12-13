package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()

	log.Printf("enter %s\n", msg)

	return func() {
		log.Printf("退出: %s (%s)\n", msg, time.Since(start))
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOpeateion")()	// 后面加括号 才能调用返回的 function value
	time.Sleep(time.Second*10)
}

func main() {
	bigSlowOperation()
}
