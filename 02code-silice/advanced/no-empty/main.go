package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.

func nonempty(strs []string) []string {

	i := 0
	for _, v := range strs {
		if v != "" {
			strs[i] = v
			i++
		}
	}

	return strs[:i]
}


func nonempty2(strs []string) []string {

	// out := strs[:0]
	var out []string

	for _, v := range strs {
		if v != "" {
			out = append(out, v)
		}
	}

	return out


}

func main() {

	data := []string{"dajl", "", "daf"}

	// fmt.Printf("nonempty(data): %q\n", nonempty(data))
	
	// fmt.Printf("data: %q\n", data)

	fmt.Printf("nonempty2(data): %q\n", nonempty2(data))

	fmt.Printf("data: %q\n", data)



}
