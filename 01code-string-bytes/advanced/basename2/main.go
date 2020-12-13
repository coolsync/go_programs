package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LastIndex returns the index of the last instance of substr in s,
// or -1 if substr is not present in s

// LastIndex 返回字符串 substr 在字符串 str 中最后出现位置的索引（ str 的第一个字符的索引），
// -1 表示字符串 sbustr 不包含字符串 str

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}


// basename removes directory components and a .suffix.
// basename删除目录组件和一个.suffix

func basename(str string) string {
	// Discard last '/' and everything before.
	// 丢弃最后的 '/' 和之前的所有内容
	slash := strings.LastIndex(str, "/")	// -1 if "/" not found

	str = str[slash+1:]

	// Preserve everything before last '.'.
	// 保留最后一个 '.' 之前的所有内容。
	
	if dot := strings.LastIndex(str, "."); dot >= 0 {
		str = str[:dot]
	}

	return str
}
