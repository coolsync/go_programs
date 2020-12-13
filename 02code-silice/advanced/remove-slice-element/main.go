package main

import "fmt"

// 移除指定index的元素
func remove(sli []int, i int) []int {
	copy(sli[i:], sli[i+1:])

	return sli[:len(sli)-1]
}

// 移除指定index的元素， 并将 最后一个元素 移到 指定的 index下标
func remove2(sli []int, i int) []int {
	sli[i] = sli[len(sli)-1]

	return sli[:len(sli)-1]
}

func main() {
	sli := []int{5, 6, 7, 8, 9}

	// fmt.Printf("%v\n", remove(sli, 2))

	fmt.Printf("%v\n", remove2(sli, 2))

}