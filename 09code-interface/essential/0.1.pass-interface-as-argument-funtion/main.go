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

type dog struct {
	age uint
}

func (d dog) breathe() {
	fmt.Println("dog breathes")
}

func (d dog) walk() {
	fmt.Println("dog walk")
}


func main() {
	l := lion{10}
	CallBreathe(l)
	CallWalk(l)

	d := dog{5}
	CallBreathe(d)
	CallWalk(d)

	// nil interface
	emptyInf("this is string")
	emptyInf(10)
	emptyInf(true)
}

// pass an interface as an arguments to a function
func CallBreathe(a animal) {
	a.breathe()
}

func CallWalk(a animal) {
	a.walk()
}

// empty interface implemnet any type
func emptyInf(a interface{}) {
	fmt.Printf("%T --> %v\n", a, a)	
}