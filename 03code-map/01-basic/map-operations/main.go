package main

import "fmt"

func main() {
	// declare
	employeeSalary := make(map[string]int)

	// adding a key value
	fmt.Println("before update: ")
	employeeSalary["tom"] = 2000
	fmt.Println(employeeSalary)

	// updating a value of key
	fmt.Println("after updata:")
	employeeSalary["tom"] = 3000
	fmt.Println(employeeSalary)

	// retrieve a value
	salary := employeeSalary["tom"]

	fmt.Printf("salary: %d\n", salary)

	// delete a key value pair
	fmt.Println("deleting key:")
	
	delete(employeeSalary, "tom")

	fmt.Println(employeeSalary)
}