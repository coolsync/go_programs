package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// get one obj html node
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks %v\n", err)
		os.Exit(1)
	}

	// get a tag href val
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	
}

func visit(links []string, n *html.Node) []string {
	// get a tag content
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 递归遍历 叶节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}