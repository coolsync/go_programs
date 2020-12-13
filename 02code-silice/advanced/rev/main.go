package main

import "fmt"
// reverse reverses a slice of ints in place.
func main() {
	a := [...]int{0,1,2,3,4,5}
	
	reverse(a[:])

	fmt.Println("a: ", a)


	s := []int{0,1,2,3,4,5}

	reverse(s[:2])	// 1 0

	reverse(s[2:])	// 5 4 3 2

	reverse(s)		// 1 0 5 4 3 2 --> 2 3 4 5 0 1

	fmt.Println("s: ", s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}