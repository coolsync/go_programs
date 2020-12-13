package main

import (
	"fmt"
	"math"
)

type employee struct {
	name string
	age uint
	salary float64
}

func (e employee) printName() employee {
	fmt.Printf("%s\n", e.name)
	return e
}

func (e employee) printAge() employee {
	fmt.Printf("%d\n", e.age)
	return e
}

func (e employee) printSalary() {
	fmt.Printf("%.2f\n", e.salary)
}

// custom type
type myFloat float64

func (m myFloat) ceil() float64 {
	return math.Ceil(float64(m))
}

func main() {
	e := employee{"sam", 31, 2000.11}

	e.printName().printAge().printSalary()

	num := myFloat(1.2)

	fmt.Println(num.ceil())
}