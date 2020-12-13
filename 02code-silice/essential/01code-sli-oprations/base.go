// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p)) // 将写入的字节数 conv type ByteCounter
	return len(p), nil
}

func main() {

	var c ByteCounter
	c.Write([]byte("hello")) // c = 5

	c.Write([]byte("yes"))
	fmt.Println(c) // 8

	c = 0 // reset c
	name := "paul"
	fmt.Fprintf(&c, "hello, %s", name) // 11
	fmt.Println(c)
}
