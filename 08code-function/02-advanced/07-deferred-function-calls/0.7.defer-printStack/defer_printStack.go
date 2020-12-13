package main

import (
	"fmt"
	"os"
	"runtime"
)

// Go’s  panic  mechanism  runs  the  deferred functions before it unwinds the stack
func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte

	n := runtime.Stack(buf[:], false)

	os.Stdout.Write(buf[:n])
}


func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	
	defer fmt.Printf("defer x=%d\n", x)

	f(x - 1)
}