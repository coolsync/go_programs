package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create map
	seen := make(map[string]bool) // bool 初始值为 false

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		
		line := input.Text()

		if line == "exit" {
			break
		}
		if !seen[line] {
			seen[line] = true
			fmt.Println(seen[line])
		} else {
			fmt.Println("已存在")
		}
	}

	if input.Err() != nil {
		fmt.Fprintf(os.Stderr, "%s", "input err")

		os.Exit(1)
	}
}
