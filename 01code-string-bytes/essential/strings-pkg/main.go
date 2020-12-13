package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	url := "https://www.bing.com"
	dir := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")
	fmt.Println(dir)

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Println(err)
	}

	filename := dir + "/index.html"
	f, _ := os.Create(filename)
	defer f.Close()
}
