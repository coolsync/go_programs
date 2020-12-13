package main

import (
	"fmt"
	"go-programs/08code-function/05-anonymous-function/links"
	"log"
	"os"
)

// 1. 创建 breadthFirst 用于所有链接及内部符合的链接
// 2. 维护一个map seen, 用于排除已添加的链接
// 3. 遍历 item, 有worklist产生,
// 4. 使用crawl 表示 function_value f, 完成一个链接内的所有链接获取

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)

	// itmes

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}	
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	link, err := links.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return link
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
