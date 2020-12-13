package main

import (
	"flag"
	"fmt"
	"go-programs/09code-interface/advanced/0.9.ex6.tempconv_kelvin/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "conv celsius temp")

func main() {
	flag.Parse() // 解析 flag
	fmt.Println(*temp)
}
