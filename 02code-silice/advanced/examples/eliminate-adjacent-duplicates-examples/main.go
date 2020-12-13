package main

// Write an in-place function to eliminate adjacent duplicates in a []string slice.
// 编写就地以消除[]字符串片段中的相邻重复项 function。

func main() {
	// s := []string{"a", "b", "b", "b", "c", "b"}
	s := []string{"a", "b", "b", "c", "b"}
	// copy(s[:], s2[:])

	// n := copy(s[4:], s[5:])
	s = remove(s)
	// fmt.Println(n, s)
}

/* Error: Unable to move 'UTF-8-encoded-spaces-examples' into 'examples' (Error: EBUSY: resource busy or locked, rename 'd:\Documents\workplay\godev\src\go-programs\02code\examples\UTF-8-encoded-spaces-examples' -> 'd:\Documents\workplay\godev\src\go-programs\02code\examples\utf8-encoded-spaces-examples'). */

func remove(s []string) []string {

	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}
