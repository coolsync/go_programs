package main

import (
	"fmt"
	"sort"
)

func main() {
	vals := []int{5, 1, 6, 1}

	// IntsAreSorted tests whether a slice of ints is sorted in increasing order.
	// IntsAreSorted测试是否按增量顺序对一个整数切片进行排序
	fmt.Println(sort.IntsAreSorted(vals)) // false

	// Ints sorts a slice of ints in increasing order.
	// 整数按递增顺序对整数切片进行排序
	sort.Ints(vals)

	fmt.Println(vals) // [1 1 5 6]

	fmt.Println(sort.IntsAreSorted(vals)) // true

	// IntSlice attaches the methods of Interface to []int, sorting in increasing order
	// IntSlice将Interface的方法附加到[] int，并按升序排序
	sort.Sort(sort.Reverse(sort.IntSlice(vals)))

	fmt.Println(vals) // [6 5 1 1]

	fmt.Println(sort.IntsAreSorted(vals)) // false
}

