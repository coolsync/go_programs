package main

import (
	"fmt"
	"reflect"
)

func main() {
	// define nil pointer int variable
	var a *int

	fmt.Printf("var a *int: %v\n", reflect.ValueOf(a).IsNil())

	// nil value
	fmt.Printf("nil value: %v\n", reflect.ValueOf(nil).IsValid())

	// *int type nil pointer
	fmt.Printf("(*int)(nil): %v\n", reflect.ValueOf((*int)(nil)).Elem().IsValid())

	// instantiate a struct
	s := struct{}{}

	fmt.Printf("get 不存在的成员： %v\n", reflect.ValueOf(s).FieldByName("").IsValid())

	fmt.Printf("get 不存在的方法： %v\n", reflect.ValueOf(s).MethodByName("").IsValid())

	// instantiate a map
	m := map[string]int{}

	fmt.Printf("get 不存在的健： %v\n", reflect.ValueOf(m).MapIndex(reflect.ValueOf("")).IsValid())
}

/* 
// Elem returns the value that the interface v contains
// or that the pointer v points to.
// It panics if v's Kind is not Interface or Ptr.
// It returns the zero Value if v is nil.
*/