package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)

	var stack []string
	for {
		tok, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			// fmt.Println("xml.EndElement b: ",stack)
			stack = stack[:len(stack)-1] // pop
			// fmt.Println("xml.EndElement a: ",stack)
			// unneeded: len(stack)simplify slice

		case xml.CharData:
			if containAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// x 是否包含 y
func containAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if y[0] == x[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
