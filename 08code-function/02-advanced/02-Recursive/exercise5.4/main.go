// Extend the visit function so that it extracts other kinds of links from the document,
// such as images, scripts, and style sheets.
// 扩展访问功能，以便从文档中提取其他类型的链接，
// 例如图片，脚本和样式表。
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var (
	links []string
	imgs []string
	scripts []string
	styles []string
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse err: %v\n", err)
		os.Exit(1)
	}

	visit(doc)

	fmt.Println(links, "\n\n")
	fmt.Println(imgs, "\n\n")
	fmt.Println(scripts, "\n\n")
	fmt.Println(styles, "\n\n")

}

func visit(n *html.Node) {
	// if n is nil, for trasever
	if n.Type != html.ElementNode {
		goto LOOP
		// for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 	visit(c)
		// }		
	}

	switch n.Data {
	case "a":
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	case "img":
		for _, s := range n.Attr {
			if s.Key == "src" {
				imgs = append(imgs, s.Val)
			}
		}
	case "script":
		for _, s := range n.Attr {
			if s.Key == "src" {
				scripts = append(scripts, s.Val)
			}
		}
	case "link":
		for _, h := range n.Attr {
			if h.Key == "href" {
				styles = append(styles, h.Val)
			}
		}
	}

LOOP: 
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}