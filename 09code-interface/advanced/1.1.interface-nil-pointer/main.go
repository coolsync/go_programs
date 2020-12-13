package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	// var buf *bytes.Buffer
	var buf io.Writer	// 正确操作
	fmt.Printf("%T, %v\n", buf, buf)

	if debug {
		buf = new(bytes.Buffer)
		fmt.Printf("%T, %#v\n", buf, buf)
	}

	f(buf)	// 传入含有空接口值的 pointer接口动态类型时, 会引发panic, 只有均为nil时, 正常进行

	fmt.Println(buf)
}

// 将 out 写入结果 传入给 buf缓冲区
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	} else {
		fmt.Println("out is nil")
	}
}
