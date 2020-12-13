package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer

	w = os.Stdout

	if f, ok := w.(*os.File); ok {
		fmt.Printf("underlying type: %T\n", f)
		fmt.Printf("underlying value: %#v\n", f)
	}
	// 声明了一个同名的新的本地变量，外层原来的w不会被改变
	if w, ok := w.(*os.File); ok {
		fmt.Printf("underlying type: %T\n", w)
		fmt.Printf("underlying value: %#v\n", w)
	}

	// c := w.(*bytes.Buffer)

	// fmt.Println(c)

	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	w = rw             // io.ReadWriter is assignable to io.Writer
	w = rw.(io.Writer) // fails only if rw == nil
	
	// The second result is conventionally assigned to a variable named ok.
	// If the operation failed, ok is false,
	// and the first result is equal to the zero value of the asserted(认定) type,
	// which in this example is a nil *bytes.Buffer.

	// 第二个结果通常分配给名为ok的变量。
	// 如果操作失败，则ok为false，
	// 并且第一个结果等于认定类型的零值，
	// 在此示例中为nil *bytes.Buffer

	if b, ok := w.(*bytes.Buffer); !ok {
		fmt.Printf("underlying type: %T\n", b)
		fmt.Printf("underlying value: %#v\n", b)
	}
}
