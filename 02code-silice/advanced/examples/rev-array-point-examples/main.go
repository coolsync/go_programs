package main

import "fmt"

// 重写reverse函数，使用数组指针代替slice

func main() {
	a := [5]int{0, 1, 2, 3, 4}

	reverse(&a)

	fmt.Printf("%v\n", a)
}

func reverse(a *[5]int) {
	for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
