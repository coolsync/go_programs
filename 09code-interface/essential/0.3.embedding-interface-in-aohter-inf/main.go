package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type human interface {
	animal
	speak()
}

type employee struct {
	name string
}

func (e employee) breathe() {
	fmt.Printf("employee %s breathes\n", e.name)
}

func (e employee) walk() {
	fmt.Printf("employee %s walk\n", e.name)
}

func (e employee) speak() {
	fmt.Printf("employee %s spokes\n", e.name)
}


func main() {
	var h human

	h = employee{"bob"}

	h.breathe()
	h.walk()
	h.speak()
}