package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type NewCustomerEvent struct {
	Name string
	Phone string
	Email string
}

func main() {
	newCustomer := NewCustomerEvent{"x", "1234779", "xya@qq.com"}

	convert(newCustomer)
}

// conversion empty interface to struct
func convert(event interface{}) {
	c := NewCustomerEvent{}

	mapstructure.Decode(event, &c)

	fmt.Printf("event is: %v\n", c)
}