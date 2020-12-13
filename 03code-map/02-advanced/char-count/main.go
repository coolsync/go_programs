package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// 1. 创建 utf_count map, 统计unicode 字符 总量
// 2. 创建 utf_len, 统计 unicode 字符 各个长度的总量
// 3. 定义 invalid, 统计 无效的 unicode 字符

// 4. use NewReader, 构建读取器
// 5. for 遍历 从 readRune 读取到的信息, rune 字符， 字节长度， error

// 6. 依次添加
// 7. 打印信息
func main() {
	// 1.
	utf_counts := make(map[rune]int)

	// 2.
	var utf_len [utf8.MaxRune + 1]int

	// 3
	invalid := 0

	// read file
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Fprintln(os.Stdin, "file read failed")
	}

	defer file.Close()

	// 4 构建读取器
	reader := bufio.NewReader(file)

	for {
		r, n, err := reader.ReadRune()

		if err == io.EOF {
			break	// 读到末尾， 跳出循环
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		utf_counts[r]++

		utf_len[n]++
	}

	fmt.Println("\nutf 字符 counts: ")

	for i, v := range utf_counts {
		fmt.Printf("\nutf8 %q \t counts: %d\n", i, v)
	}

	fmt.Println("\nutf 字符 各个固定长度 counts: ")

	for i, v := range utf_len {
		if v > 0 {

			fmt.Printf("\nutf len  %v \t counts: %d\n", i, v)
		}
	}

	fmt.Printf("\n无效的字符个数： %d\n", invalid)
}
