package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 从 os.Args 获取 url
// title(), get resp pkg, parse resp
// 使用 resp.Url.Header(), get Content-Type内容, 无 text/html, 返回错误
// visitAll(), 用于具体节点操作, get title node 文本内容
// forEachNode, 用于遍历节点操作

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 获取 Content-Type
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("url %s, content-type err, not text html: %v\n", url, err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("url %s, parse as html err: %v\n", url, err)
	}

	visitAll := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data) // 文本节点
		}
	}

	forEachNode(doc, visitAll, nil)

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
		if err := title(url); err != nil {
			fmt.Fprintf(os.Stderr, "title err: %v\n", err)
		}
	}
}
