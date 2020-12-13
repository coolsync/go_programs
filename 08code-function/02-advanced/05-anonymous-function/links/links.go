package links

import (
	"fmt"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

// 1. 创建 Extract function, 提取有效 a link
// 2. 获取url response package, 并使用html包解析resp
// 3. 创建 function value visitAll, 具体完成节点查找操作
// 4. 使用 forEachNode, 完成节点遍历操作
// 5. 将结果返回并展示

func Extract(url string) ([]string, error) {
	// 获取resp包
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	// 判断状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s : %d\n", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("pasing %s as html, err: %v\n", url, err)
	}

	var links []string
	// 创建visitall 功能, 用于节点操作
	visitAll := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {

				if a.Key != "href" {
					continue
				}
				// links = append(links, a.Val)
				// 解析 url格式是否符合规范
				if !strings.HasPrefix( a.Val, "http") || !strings.HasPrefix(a.Val, "https") {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // 忽视解析错误
				}
				links = append(links, link.String())
			}
		}

	}

	forEachNode(doc, visitAll, nil)

	// resp.Close = true
	return links, nil
}

// 遍历node 操作
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

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Fprintln(os.Stderr, "usage: xx.go url")
// 		os.Exit(1)
// 	}

// 	url := os.Args[1]

// 	links, err := Extract(url)
// 	if err != nil {
// 		log.Fatal("extract err: ", err)
// 	}
// 	for i, link := range links {
// 		fmt.Printf("%d\t%s\n", i, link)
// 	}
// }
