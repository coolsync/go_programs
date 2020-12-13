package main

import "fmt"

// defer FIOT first in out to
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)	// panics if x == 0

	defer fmt.Printf("defer f(%d)\n", x)

	f(x - 1)
}

func main() {
	f(3)
}