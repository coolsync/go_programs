package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age uint
}

func (l lion) breathe() {
	fmt.Println("lion breathes")
}

func (l lion) walk() {
	fmt.Println("lion walk")
}

func main() {
	var a animal
	var b animal
	var c animal
	var d animal
	var e animal

	a = lion{10}
	b = lion{10}
	c = lion{5}

	if a == b {
		fmt.Println("a and b is equal")
	} else {
		fmt.Println("a and b is not equal")
	}

	if a == c {
		fmt.Println("a and c is equal")
	} else {
		fmt.Println("a and c is not equal")
	}

	if d == e {
		fmt.Println("d and e is equal")
	} else {
		fmt.Println("d and e is not equal")
	}
}