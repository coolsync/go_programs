package main

import "fmt"

// 需求:  Use panic and recover to write a function
// that contains no return statement yet returns a non-zero value

func main() {
	fmt.Println(returnNoZero(6))
}

func returnNoZero(x int) (result int) {
	defer func() {
		_ = recover()	// recover后 自动调用 return
		result = x+1
	}()

	panic("panic!")
}