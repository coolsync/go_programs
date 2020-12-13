package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string
	Age  int
}

func main() {
	var h Human

	// get reflect.Type
	typeOfh := reflect.TypeOf(h)

	fmt.Printf("name: %v\nkind: %v\n", typeOfh.Name(), typeOfh.Kind())
}
