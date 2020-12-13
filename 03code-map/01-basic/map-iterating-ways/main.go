package main

import "fmt"

func main() {
	// define a map
	simple := map[string]string{
		"a": "x",
		"b": "y",
	}

	// iterating all keys and values
	for key, value := range simple {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	// iterating only keys

	for key := range simple {
		fmt.Printf("key: %s\n", key)
	}
	
	keys := getAllKeys(simple)

	fmt.Println(keys)
}

// get list of all keys
func getAllKeys(simple map[string]string) []string {
	
	var keys []string
	for k := range simple {
		keys = append(keys, k)
	}

	return keys
}
