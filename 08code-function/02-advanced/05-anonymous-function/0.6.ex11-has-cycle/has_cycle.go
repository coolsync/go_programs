package main

import (
	"fmt"
	"log"
	"sort"
)

// 1. 创建 toposort, 完成提取有向图的任务,将结果返回给 sorted
// 2. order []string, 用于保存递归提取元素
// 3. seen map, 用于消除已提取的元素
// 4. visitAll([]string), 递归提取元素
// 5. 排序课程的key值, 即先决条件
// 6. 传入key值 到visitAll作为参数

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
	"linear algebra":        {"calculus"},
}

func toposort(m map[string][]string) ([]string, error) {
	// order 保存提取后的元素
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
			} else {
				hasCycle := true

				for _, s := range order {
					if s == item {
						hasCycle = false
					}
				}
				if hasCycle {
					return fmt.Errorf("存在 闭环 %s\n", item)
				}
			}
		}

		return nil
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	if err := visitAll(keys); err != nil {
		return nil, err
	}

	return order, nil
}

func main() {
	sorted, err := toposort(prereqs)
	if err != nil {
		log.Fatal(err)
	}

	for i, course := range sorted {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}
