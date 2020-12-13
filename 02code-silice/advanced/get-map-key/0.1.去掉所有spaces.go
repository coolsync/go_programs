package main

import (
	"fmt"
)

func main() {

	sample := map[string]string {
		"a": "x",
		"b": "y",
	}

	res := getKey(sample)

	fmt.Println("res: ", res)
}

func getKey(sample map[string]string) []string {
	var keys []string

	for k := range sample {
		keys = append(keys, k)
	}

	return keys
}
