package main

import (
	"flag"
	"fmt"
	"time"
)

// flag.Duration("param name", default value, "help info")

// 创建标志位param
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	// 解析flag
	flag.Parse()

	fmt.Printf("sleeping... %v\n", *period)
	time.Sleep(*period)
	
	fmt.Println()
}