package main

import (
	"bytes"
	"fmt"
	"os"
)

// 需求： 一个 非递归版本的comma 功能，使用bytes.Buffer代替 字符串链接操作

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(comma(os.Args[i]))
	}
}

func comma(str string) string {
	n := len(str)

	if n <= 3 {
		return str
	}

	var buf bytes.Buffer

	for i, v := range str {
		if i%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf,"%c", v)
	}

	return buf.String()
}
