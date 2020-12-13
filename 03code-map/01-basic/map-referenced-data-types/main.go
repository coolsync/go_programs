package main

import "fmt"


func main() {
	// declare
	employeeSalary := make(map[string]int)

	// adding a key value
	employeeSalary["tom"] = 20000

	employeeSalary["bob"] = 10000
	
	es := employeeSalary

	// change employeeSalary
	employeeSalary["alice"] = 30000

	fmt.Println("changing employeeSalary: ", employeeSalary)
	fmt.Println("es: ", es)

	// change es

	es["alice"] = 15000
	fmt.Println("changing es: ", employeeSalary)
	fmt.Println("employeeSalary: ", employeeSalary)

}