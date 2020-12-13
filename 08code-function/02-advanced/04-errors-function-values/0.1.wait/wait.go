package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// 生成 deadline 截止时间
// 遍历 deadline 之前的每一个点
// 发起 重定向
// 判断 err 是否 为 nil
// 打印每个点过后 错误信息
// 指数级增加 tries 睡眠后， 再进行connect
// 打印 deadline 后的 错误信息

func waitForServer(url string) error {
	const timeout = 1 * time.Minute

	// 生成 deadline 截止时间
	deadline := time.Now().Add(timeout)

	// 遍历 deadline 之前的 每一个点
	for tries := 0; time.Now().Before(deadline); tries++ {
		// 判断 重定向
		_, err := http.Head(url)
		// nil 直接返回 nil
		if err == nil {
			return nil
		}
		// 打印 每个点的执行后 错误的日志信息
		log.Printf("server not respoding(%s), retrying...\n", err)
		// 指数级 增加 tries 睡眠后， 再进行connect
		fmt.Printf("after %v tries\n", time.Second << uint(tries))
		time.Sleep(time.Second << uint(tries))
	}
	// 返回 deadline后错误信息 
	return fmt.Errorf("server %s failed after %s\n", url, timeout)
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Usage: wait url")
	}
	
	url := os.Args[1]

	if err := waitForServer(url); err != nil {
		log.Fatalf("waitForServer function: %s\n", err)
	}
}

