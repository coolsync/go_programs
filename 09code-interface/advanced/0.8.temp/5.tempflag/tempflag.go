package main

import (
	"flag"
	"fmt"
	tempconv "go-programs/09code-interface/advanced/0.8.temp/4.tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse() // 解析 flag 标志位
	fmt.Println(*temp)
}
