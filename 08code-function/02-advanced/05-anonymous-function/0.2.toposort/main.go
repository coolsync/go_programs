package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func toposort(m map[string][]string) []string {
	// 找出 prereqs key值, 进行排序
	// 将key值的map放入 function visitAll
	// 如果已遍历, 则使seen[item]为true, 不用再递归遍历
	// 将结果存入 order 切片

	var order []string 
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				fmt.Println("m[item] ",m[item])
				visitAll(m[item])
				order = append(order, item)
				fmt.Println("order", order)
			}
		}
	}
	
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	fmt.Printf("keys: %q\n", keys)
	sort.Strings(keys)

	visitAll(keys)
	// for _, key := range keys {
	// 	visitAll(m[key])
	// }

	return order

}

func main() {
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}
