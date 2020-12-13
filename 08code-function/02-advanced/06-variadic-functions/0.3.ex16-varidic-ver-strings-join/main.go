package main

import (
	"fmt"
	"strings"
)

// 需求: Write a variadic version of strings.Join
// strings.Join()

/*
	func Join(elems []string, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	var b Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

*/

func Join(elems []string, sep string) string {
	// 只有0个或一个元素的切片
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[1]
	}

	// 总共有多少个字符
	n := len(sep) * (len(elems)-1)

	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	// 使用 Buffer 写入字符
	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}


func JoinVariadic(sep string, elems ...string ) string {
	// sep := elems[len(elems)-1]
	// elems = elems[:len(elems)-1]


	// 只有0个或一个元素的切片
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[1]
	}

	// 总共有多少个字符
	n := len(sep) * (len(elems)-1)

	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	// 使用 Buffer 写入字符
	var b strings.Builder
	b.Grow(n)
	// b.WriteString(elems[0])
	// for i := range elems[0:] {
	// 	if i == len(elems)-1 {
	// 		b.WriteString(elems[i])
	// 		break
	// 	}
	// 	b.WriteString(elems[i])
	// 	b.WriteString(sep)
	// 	fmt.Println(elems[i])
	// }

	b.WriteString(elems[0])
	for _,s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}
func main() {
	sli := []string{"abc", "123", "efg", "fadf"}
	str := JoinVariadic("  ", sli...)

	fmt.Printf("%q\n", str)
}
