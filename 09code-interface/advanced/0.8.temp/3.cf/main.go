package main

import (
	"fmt"
	tempconv "go-programs/09code-interface/advanced/0.8.temp/2.tempconv"
	"os"
	"strconv"
)

// Cf converts its numeric argument to Celsius and Fahrenheit.
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)	// 字符串 转换 float64
		if err != nil {
			fmt.Fprintf(os.Stderr, "fromat float err: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t) // 转换类型
		f := tempconv.Fahrenheit(t)

		fmt.Printf("%s=%s,\t%s=%s\n", c, tempconv.C_Conv_F(c), f, tempconv.F_Conv_C(f))
	}
}

