package main

import "fmt"

func main() {

	// Write a version of rotate 旋转 that operates in a single pass.
	// 在指定的位置上 Rotate 原切片。

	s := []int{0, 1, 2, 3, 4}
	

	// n := 2 % 5

	// tmp := append(s, s[:n]...)

	// fmt.Println(tmp[n:])
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, n int) {
	// n %= len(s)
	n = n % len(s)
	tmp := append(s, s[:n]...)
	copy(s, tmp[n:])
}
