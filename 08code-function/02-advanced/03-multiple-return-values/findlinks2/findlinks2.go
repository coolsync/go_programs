// Findlinks2 does an HTTP GET on each URL, parses the
// result as HTML, and prints the links within it.

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinsk2: %v\n", err)
			continue
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.

func findlinks(url string) ([]string, error) {
	// get resp pkg
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// judge status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting url %s failed\n", url)
	}

	// parse html node
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing html %s, err: %v\n", url, err)
	}

	// call visit finished traverse html node
	return visit(nil, doc), nil
}

// visit appends to links each link found in n, and returns the result 
func visit(links []string, n *html.Node) []string {
	// find a tag href attr value
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// traverse leaf node and sibling node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}