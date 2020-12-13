package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

// 封装 io.Writer
func CountingWirter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{w, 0}
	return &c, &c.writtern
}

type ByteCounter struct {
	w        io.Writer
	writtern int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.writtern += int64(n)
	return n, err
}

func main() {
	str := "hello paul!"

	w, n := CountingWirter(ioutil.Discard) // get 返回值输入
	fmt.Fprintf(w, str)                    // 字符串 转为 字节
	fmt.Println(*n)                        // 11
}
