package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3, 4)) // 先将其变为array, 然后在转为切片, 最后传入到sum(), 作为可变params
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)


}

func f(...int) {}
func g([]int)  {}

func sum(vals ...int) int {
	total := 0

	for _, val := range vals {
		total += val
	}

	return total
}

func f1() {
	
	linenum, name := 12, "count"
	errorf(linenum, "未定义: %s", name)

}

// Variable parameter functions are often used to format strings
func errorf(linenum int, format string, arg ...interface{}) {
	fmt.Fprintf(os.Stderr, "line: %d", linenum)
	fmt.Fprintf(os.Stderr, format, arg...)
	fmt.Fprintln(os.Stderr)
}
