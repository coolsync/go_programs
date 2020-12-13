package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Write a function to populate a mapping from element names—p, div, span,
// and so on—to the number of elements with that name in an HTML document tree.

// 创建一个功能，记录在HTML树中出现的同名元素的次数
var count = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findElements: %v\n", err)
	}
	

	visit(doc)

	for t, c := range count {
		fmt.Println(t, c)
	}
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}