package main

import (
	"bytes"
	"fmt"
	"os"
)

// 编一个function，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序
func main() {

	// str1 := "abcd"

	// str2 := "bcda"

	b := compareStr(os.Args[1], os.Args[2])



	fmt.Println("str1 和 st2 是否包含相同字符： ", b)

}

// method 1
func compareStr(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	m := make(map[int][]byte)

	var buf1 bytes.Buffer
	var buf2 bytes.Buffer

	// 提取字符
	for _, v := range str1 {
		fmt.Fprintf(&buf1, "%c", v)
		// buf1.WriteByte(v)

	}

	m[1] = buf1.Bytes()

	for _, v := range str2 {
		fmt.Fprintf(&buf2, "%c", v)
	}

	m[2] = buf2.Bytes()


	for _, v := range m[1] {
		if !bytes.Contains(m[2], []byte{v}) {
			return false
		}
	}

	return true
}


// method 2

func compareStr2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	m := make(map[rune]int, len(str1))

	for _, v := range str1 {
		m[v]++
	}

	for _, v := range str2 {
		if m[v] == 0 {
			return false
		}

		m[v]--
	}

	return true
}