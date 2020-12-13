package main

import "fmt"

func main() {
	// declare
	employeeSalary := make(map[string]int)


	// adding a key value 
	employeeSalary["tom"] = 20000

	fmt.Println("key 存在 map: ")

	val, ok := employeeSalary["tom"]
	fmt.Printf("val %d, ok: %t\n", val, ok)


	fmt.Println("key 不存在 map: ")
	
	val, ok = employeeSalary["sam"]
	fmt.Printf("val %d, ok: %t\n", val, ok)
}
	