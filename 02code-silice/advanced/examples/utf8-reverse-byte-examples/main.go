package main

import (
	"fmt"
	"unicode/utf8"
)

// 修改reverse功能 用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

func main() {
	b := []byte("中 国 字")

	reverse_utf8(b)

	fmt.Printf("%q\n", b)
}

func reverse_utf8(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		// fmt.Println(b)

		reverse(b[i : i+size])

		i += size
	}

	reverse(b)
}

func reverse(b []byte) {
	last := len(b) - 1

	for i := 0; i < len(b)/2; i++ {
		b[i], b[last-i] = b[last-i], b[i]
	}
	// fmt.Println(b)
}
