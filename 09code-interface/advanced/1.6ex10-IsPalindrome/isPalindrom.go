package main

import (
	"fmt"
	"sort"
)

func main() {
	// str := "abccba" // true 首尾一致
	str := "aebceg" // false

	// []byte sort
	b := bs(str)
	sort.Sort(b)
	fmt.Println(b)

	fmt.Println("是否有回文: ", isPalindrom(bs(str)))
}

type bs []byte

// 实现sort.Interface
func (b bs) Len() int {
	return len(b)
}

func (b bs) Less(i, j int) bool {
	return b[i] < b[j] // 前一个元素小于后一个元素, index不变, 交换值
}

func (b bs) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// 判断 slice 首尾元素是否有一致内容
func isPalindrom(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
