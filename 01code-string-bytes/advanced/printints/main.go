package main

import (
	"bytes"
	"fmt"
)

// Printints demonstrates the use of bytes.Buffer to format a string.

// 需求： 使用 bytes.Buffer 将 [1,2,3] 格式化 字符串

// intsToString is like fmt.Sprint(values) but adds commas.
// intsToString 类似于 fmt.Sprintf（values），区别是 添加逗号。 

func intsToString(values []int) string {
	var buf bytes.Buffer

	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')

	return buf.String()
}

func main(){
	fmt.Printf("%q", intsToString([]int{1,2,3}))
}