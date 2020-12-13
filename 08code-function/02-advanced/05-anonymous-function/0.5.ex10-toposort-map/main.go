package main

import "fmt"

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"liniear algebra": true},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func toposort(m map[string]map[string]bool) []string {
	var order []string
	// seen map, 用于消除已查看的元素
	seen := make(map[string]bool)
	
	// 创建 func_val visitAll, 用于递归遍历子元素, 添加到 order内
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for k := range items {
			if !seen[k] {
				seen[k] = true
				visitAll(m[k])
				order = append(order, k)
			}
		}
	}

	// 遍历 prereqs, 获取key值
	for k := range m {
		visitAll(map[string]bool{k: true})
	}
	// 将order结果返回, 并遍历
	return order
}

func main() {
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

