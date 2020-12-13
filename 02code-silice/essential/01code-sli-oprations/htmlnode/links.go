package links

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// func main() {
// 	for _, url := range os.Args[1:] {
// 		links, err := Extract(url)
// 		if err != nil {
// 			log.Fatalf("Extract err: %v", err)
// 		}
// 		for i := range links {
// 			fmt.Printf("%d\t%s\n", i, links[i])
// 		}
// 	}

// }

// Extract 实现了 解析 html 和 操作节点 的两个作用
func Extract(url string) ([]string, error) {
	// 获取应答包
	resp, err := http.Get(url)
	if err != nil {
		if err == io.EOF {
			goto Loop
		}
		return nil, fmt.Errorf("http get: %v", err)
	}
Loop:
	defer resp.Body.Close()

	// 判断应答包 status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// 解析 html
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http parse: %v", err)
	}

	// links []string 有效 link 的集合
	var links []string

	// 创建闭包 完成 操作节点
	var visit = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {

					if strings.HasPrefix(a.Val, "javascript") {
						a.Val = ""
					}

					// 判断 link 是否 是 有效链接
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue // ignore bad link
					}

					links = append(links, link.String())
				}
			}
		}
	}

	forEachNode(doc, visit, nil)

	return links, nil
}

// 递归遍历node
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil { // 遍历 child node before
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil { // 遍历 child node after
		post(n)
	}
}
