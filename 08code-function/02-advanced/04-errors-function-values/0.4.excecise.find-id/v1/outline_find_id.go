package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "http://www.qq.com"

	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	check(err)
	

	forEachElement(doc)
}

func forEachElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// fmt.Println(n.Data)
		for _, a := range n.Attr {
			fmt.Printf("key: %T\t%v\n", a.Key, a.Key)
			fmt.Printf("value: %T\t%v\n", a.Val, a.Val)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachElement(c)
	}
}
