package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}

	// fmt.Println(basename("b/c/a.go"))
}

func basename(str string) string {
	// b/c
	// 提取 最后一个 / 后的内容
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '/' {
			str = str[i+1:]
			break
		}
	}

	// 提取 . 前的内容
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '.' {
			str = str[:i]
			break
		}
	}
	return str
}
// runtime error: index out of range [4] with length 4
