package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outlink err: %v\n", err)
		os.Exit(1)
	}

	outlink(nil, doc)
}

func outlink(stack []string,n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)	// push tag in stack
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outlink(stack, c)
	}
}