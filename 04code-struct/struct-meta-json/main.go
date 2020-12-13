package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// type employee struct {
// 	Name   string `json:"name"`
// 	Age    int    `json:"age"`
// 	Salary int    `json:"salary"`
// }

type employee struct {
	Name   string 
	Age    int    
	Salary int    
}


func main() {
	emp := employee{Name: "小黑", Age: 18, Salary: 3000}

	// accesing a struct field
	fmt.Printf("%#v\n", emp)

	// assigning a new value
	emp.Name = "小红"
	fmt.Printf("%#v\n", emp)

	data, err := json.Marshal(&emp)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Marshal function output: %s\n", data)

	jsonEmp, err := json.MarshalIndent(&emp, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MarshalIndent function output: \n%s\n", jsonEmp)
}
