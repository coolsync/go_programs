package main

import "fmt"

type employee struct {
	name string
	age uint
	salary float64
	address
}

type address struct {
	country, city string
}

func (emp employee) details() {
	fmt.Printf("Name: %s\n", emp.name)
	fmt.Printf("age: %d\n", emp.age)
}

func (a address) addrDetails() {
	fmt.Printf("country: %s\n", a.country)
	fmt.Printf("city: %s\n", a.city)
}

func (emp employee) getSalary() float64 {
	return emp.salary
}

func (emp *employee) setNewName(newName string) {
	emp.name = newName
}

func main() {
	emp := employee{"sam", 30, 2000.11, address{"chian", "guangzhou"}}

	emp.addrDetails()

	emp.address.addrDetails()
	// emp.details()
	// fmt.Println("field name of the receiver changed inside the method:")
	// (&emp).setNewName("bob")
	// emp.details()
	// fmt.Printf("salary: %.2f\n", emp.getSalary())
}
