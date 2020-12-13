package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		// fmt.Printf(" %s\n", os.Args[i])
		
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}

// 在不是负数的整数 之间插入 逗号

func comma(str string) string {

	n := len(str)

	if n <= 3 {
		return str
	}

	return comma(str[:n-3]) + "," + str[n-3:] // :n-3 	index 0,1,2
}
