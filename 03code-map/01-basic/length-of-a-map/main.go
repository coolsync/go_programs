package main

import "fmt"

func main() {
	// declare
	// var employeeSalary map[string]int		// nil map


	employeeSalary := map[string]int{}
	fmt.Println(employeeSalary)	

	// adding a key value

	employeeSalary["tom"] = 2000
	employeeSalary["sam"] = 1200

	lenOfMap := len(employeeSalary)

	fmt.Println(lenOfMap)
}

