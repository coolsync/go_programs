package main

import "fmt"

func main() {
	// cp a string to a slice of bytes
	src := "abc"
	dst := make([]byte, 3)

	numberOfElementsCopied := copy(dst, src)
	fmt.Printf("number of elements of copied: %d\n", numberOfElementsCopied)

	fmt.Printf("src %v\n", src)
	fmt.Printf("dst %v\n", dst)
	
	fmt.Println("----------------------")
	
	srcSli := []int{1,2,3,4,5}
	numberOfElementsCopied2 := copy(srcSli, srcSli[3:])
	
	fmt.Printf("number of elements of copied: %d\n", numberOfElementsCopied2)
	
	fmt.Printf("srcSli: %v\n", srcSli)
	
}