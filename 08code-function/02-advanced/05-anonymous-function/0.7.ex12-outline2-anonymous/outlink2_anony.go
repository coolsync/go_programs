package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 需求: 解析html, 在元素前后添加格式
// 1. os.Args 获取url
// 2. outlink(url), 用于获取并解析response pkg
// 3. startElement(*html.Node) 匿名func, 用于child node之前单个节点具体操作
// 4. endElement(*html.Node) 匿名func, 用于child node之后节点具体操作
// 5. forEachNode 完成节点的遍历操作

var depth int

func outlink(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
		}
	}
	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		if err := outlink(url); err != nil {
			log.Fatal(err)
		}
	}
}
