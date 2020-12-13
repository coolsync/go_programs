package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordsCounter int // count 文本 words

func (wc *WordsCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*wc++
	}

	return len(p), nil
}

type LineCounter int // count 文本 lines

func (lc *LineCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p)) // 创建扫描器

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*lc++
	}

	return len(p), nil
}

func main() {
	str := "Hello, World\nHello, 世界"

	var wc WordsCounter

	fmt.Fprintf(&wc, str) // 返回每次读取的字节数,err
	fmt.Println(wc)       // 4

	var lc LineCounter

	fmt.Fprintf(&lc, str) // 将字符串转为 字节slice
	fmt.Println(lc)       // 2
}
