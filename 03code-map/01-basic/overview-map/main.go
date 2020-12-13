package main

import (
	"fmt"
)

func main() {
	// declare
	employeeSalary := map[string]int{}

	fmt.Println(employeeSalary)

	// intialize using map lieteral
	employeeSalary = map[string]int{
		"Alice": 2000,
		"Tom": 3000,
	}

	// adding a key value
	employeeSalary["Sam"] = 1500

	fmt.Println(employeeSalary)
}
