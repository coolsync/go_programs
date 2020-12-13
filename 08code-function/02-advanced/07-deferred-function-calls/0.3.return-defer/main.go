package main

import "fmt"

// The function in the defer statement will be executed
// after the return statement updates the return value variable

func double(x int) (result int) {
	defer func(){
		fmt.Printf("double(%d), result: %d\n", x, result)	// defer 在 return 更新后执行
	}()
	return x + x
}

func triple(x int) (result int) {
	defer func() {
		result += 4
		// fmt.Printf("tripe(%d), result: %d\n", x, result)
	}()

	return double(4)
}
func main() {
	// _ = double(4)
	fmt.Println(triple(4))
}