package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// pass url arg
// extents resposen, get html.Node
// get words imgs
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s, url", os.Args[0])
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {

		words, images, err := CountsOfWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountsOfWordsAndImages: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("words\t%d\timages%d\n", words, images)
	}
}

func CountsOfWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		// return 0, 0, err
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		// return 0, 0, err
		return
	}
	words, images = countsOfWordsAndImages(doc)
	return
}

func countsOfWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	// judge if not is textnode
	// use scanner 提取 words
	if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))

		input.Split(bufio.ScanWords)

		for input.Scan() {
			// fmt.Println(input.Text())
			words++
		}

		// words += len(strings.Fields(n.Data))
	}

	// img node,
	if n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wc, ic := countsOfWordsAndImages(c)
		words, images = words+wc, images+ic
	}

	return
}
