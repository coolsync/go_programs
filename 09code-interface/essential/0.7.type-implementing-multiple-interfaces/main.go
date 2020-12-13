package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type mammal interface {
	feed()
}

type lion struct {
	age uint
}

// lion implementing animal all method
func (l lion) breathe() {
	fmt.Println("lion breathe")
}

func (l lion) walk() {
	fmt.Println("lion walk")
}

// lion implementing mammal all method
func (l lion) feed() {
	fmt.Println("lion feed young")
}

// non-struct custom type implementing animal
// type dog string
type dog struct {
	age uint
}

func (d dog) breathe() {
	fmt.Printf("dog %T, %v breathe\n", d, d)
}
func (d dog) walk() {
	fmt.Println("dog walk")
}

func main() {
	l := lion{10}

	var a animal
	a = l
	print(a)
	fmt.Printf("underlying type: %T, underlying value: %v\n", a, a)
	a.breathe()
	a.walk()
	fmt.Println("---------------------")
	var m mammal
	m = l
	m.feed()

	// non-struct costom type
	// a = dog("wanwan")
	// a.breathe()
	// a.walk()
	a = dog{5}
	print(a)
}

func print(a animal) {
	// if l, ok := a.(lion); ok {
	// 	fmt.Println(l)
	// 	fmt.Printf("Age %d\n", l.age)
	// }else {
	// 	fmt.Println("a is not of type lion")

	// }
	// if d, ok := a.(dog); ok {
	// 	fmt.Println(d)
	// 	fmt.Printf("Age %d\n", d.age)
	// } else {
	// 	fmt.Println("a is not of type dog")
	// }

	switch v := a.(type) {
	case lion:
		fmt.Println("type lion")
	case dog:
		fmt.Println("type dog")
	default:
		fmt.Println("unknown type: ", v)
	}
	
}