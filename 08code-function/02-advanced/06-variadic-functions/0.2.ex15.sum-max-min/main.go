package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(max(7, 2, 9, 4, 9))
	fmt.Println(min(7, 2, 9, 4, 9))

	fmt.Println(max2(1))
	fmt.Println(min2())

}

func max(vals ...int) int {
	maxNum := vals[:1][0]
	for _, v := range vals {
		if maxNum < v {
			maxNum = v
		}
	}

	return maxNum
}

func min(vals ...int) int {
	// minNum := 0
	minNum := vals[:1][0]
	for _, v := range vals {
		if minNum > v {
			minNum = v
			// fmt.Println(minNum)
		}
	}

	return minNum
}

func max2(vals ...int) int {
	if len(vals) == 0 {
		log.Fatal("至少一个parameter")
	}
	maxNum := vals[:1][0]
	for _, v := range vals {
		if maxNum < v {
			maxNum = v
		}
	}

	return maxNum
}

func min2(vals ...int) int {
	if len(vals) == 0 {
		log.Fatal("至少一个parameter")
	}
	// minNum := 0
	minNum := vals[:1][0]
	for _, v := range vals {
		if minNum > v {
			minNum = v
			// fmt.Println(minNum)
		}
	}

	return minNum
}
