package main

import (
	"fmt"
	"os"
)

/* 
package os

// PathError records an error and the operation and file path that caused it.
// PathError 记录错误 以及 导致错误的 操作和文件路径
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}	
*/

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err)	// open /no/such/file: The system cannot find the path specified.

	fmt.Printf("%#v\n", err)	// &os.PathError{Op:"open", Path:"/no/such/file", Err:0x3}

	fmt.Println(os.IsNotExist(err))	// true
}