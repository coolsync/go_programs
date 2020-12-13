package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	fmt.Printf("%T\t%v\n", w, w)

	w = os.Stdout
	fmt.Printf("%T\t%#+v\n", w, w)

	w.Write([]byte("hello\n"))
	fmt.Printf("%T\t%v\n", w, w)


	w = new(bytes.Buffer)
	fmt.Printf("%T\t%#+v\n", w, w)

	w.Write([]byte("hello\n"))
	
	fmt.Printf("%T\t%v\n", w, w)


	fmt.Println("-------------")

	var x interface{} = time.Now()
	fmt.Printf("%T\t%#v\n", x, x)

}