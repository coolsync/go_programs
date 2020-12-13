package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// send req, get res
	if len(os.Args) != 2 {
		log.Fatal("Usage: outline url")
	}

	url := os.Args[1]
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	check(err)
	
	forEachElement(doc, startElement, endElement)
}


func forEachElement(n *html.Node, pre, post func(x *html.Node)) {
	if pre != nil {
		pre(n)
	}
	
	for c := n.FirstChild;  c != nil; c = c.NextSibling {
		forEachElement(c, pre, post)
	}
	
	if post != nil {
		post(n)
	}
}

var deepth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", deepth*4, "", n.Data)
		deepth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		deepth--
		fmt.Printf("%*s<%s/>\n", deepth*4, "", n.Data)
	}
}
