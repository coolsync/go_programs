// Write a function to print the contents of all text nodes in an HTML document tree.
// Do not descend into <script> or <style> elements,
// since their contents are not visible in a web browser.

//写一个功能 以打印HTML文档树中所有文本节点的内容。
//不要归入<script>或<style>元素，
//因为它们的内容在Web浏览器中不可见。
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse err: %v\n", err)
	}

	for _, texts := range visit(nil, doc) {
		fmt.Println(texts)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		for _, text := range n.Attr {
			texts = append(texts, text.Val)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling{
		texts = visit(texts, c)
	}

	return texts
}