package main

import "fmt"

func appendSlice(x []int, num ...int) []int {
	var z []int 

	z_len := len(x) + len(num)

	if z_len <= cap(x) {
		z = x[:z_len]
	} else {
		z_cap := z_len

		if z_cap < 2 * len(x) {
			z_cap = 2 *len(x)
		}

		z = make([]int, z_len, z_cap)
		
		copy(z, x)
	}

	copy(z[len(x):], num)

	return z
}


func appendInt(x []int, num int) []int {
	var z []int

	z_len := len(x) + 1

	if z_len <= cap(x) {
		// There is room to grow. Extend the slice.
		//有成长的空间。扩展切片。
		z = x[:z_len]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		//空间不足。分配一个新的数组。
		//通过加倍增长，以实现摊销的线性复杂度

		z_cap := z_len

		if z_cap < 2 * len(x) {
			z_cap = 2 * len(x)
		}

		z = make([]int, z_len, z_cap)

		copy(z, x)
	}

	z[len(x)] = num

	return z
}

func main() {
	
	var x, y []int
	
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)

		x = y
	}
	
	nums := []int{10,11,12}
	res := appendSlice(x, nums...)
	
	fmt.Println(res)

	// var runes []rune

	// for _, v := range "holo, 星空" {
	// 	runes = append(runes, v)
	// }

	// fmt.Printf("%q\n", runes)

	
}
