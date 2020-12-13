package main

import "fmt"

type animal interface {
	breathe()
	walk()
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

// If the embedded interface is a named field,
//  then interface methods have to be called via the named interface name
type pet1 struct {
	name string
	a animal
}

// If the embedded interface is unnamed/anonymous field 
// then interface methods can be referred directly or via the interface name
type pet2 struct {
	name string
	animal
}

func main() {
	d := dog{3}
	p1 := pet1{name:"pet1", a: d}
	fmt.Println("p1 name: ", p1.name)
	p1.a.breathe()
	p1.a.walk()
	// p1.breathe()		// embedded interface as named filed, have to via the named interface name
	// p1.walk()

	fmt.Println("--------------------------")

	fmt.Println("p1 name: ", p1.name)
	p2 := pet2{name:"pet2", animal: d}
	p2.animal.breathe()	// unnamed/anonymous filed, can be referred directly or via the interface name
	p2.animal.walk()
	p2.breathe()
	p2.walk()

}