package main

import (
	"fmt"
	"sort"
)

func main(){
	// ages := make(map[string]int)

	ages := map[string]int{
		"tom": 25,
		"bob": 27,
		"alice": 30,
	}
	
	// create a slice,
	names := make([]string, 0)

	// iterator map
	for name := range ages {
		names = append(names, name)
	}

	// sort names
	sort.Strings(names)

	// iterator names 
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}