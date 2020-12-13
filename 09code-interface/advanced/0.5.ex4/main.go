package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type myReader struct {
	s string
	i int64
}

func (r *myReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}

	n = copy(b, r.s[r.i:])

	r.i += int64(n)

	return n, nil
}

func NewReader(s string) io.Reader {
	return &myReader{s, 0}
}

func main() {

	s := "<h1>tag h1</h1>"

	r := NewReader(s)
	fmt.Println(r)

	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse %s err: \n", err)
		os.Exit(1)
	}

	outlink(nil, doc)
}

func outlink(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outlink(stack, c)
	}
}
