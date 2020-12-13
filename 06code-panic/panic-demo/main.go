package main

import "fmt"

// 几种 panic 演示
func main(){
	// nil pointer
	// panic: runtime error: invalid memory address or nil pointer dereference
	// var intPtr *int
	// fmt.Println(*intPtr)

	// slice
	// panic: runtime error: index out of range [10] with length 5
	// sli := make([]int, 0)	
	// sli = append(sli, 1,2,3,4,5)
	// fmt.Println(sli[10])

	// map
	// panic: assignment to entry in nil map	// 恐慌：分配给 nil map 的条目 
	var m map[string]string
	m["name"] = "jana"
	fmt.Println(m)
}