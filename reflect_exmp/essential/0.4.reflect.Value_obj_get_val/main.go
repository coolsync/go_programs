package main

import (
	"fmt"
	"reflect"
)

func main() {
	// define int type, and assign a value
	var a int = 1024
	
	// by reflect.Value create reflect value object
	valueOfA := reflect.ValueOf(a)
	
	// 通过 interface{}, type assertion get val
	getA := valueOfA.Interface().(int)

	// 通过 Int method  get val, 并将其转为int类型
	getB := int(valueOfA.Int())
	fmt.Println(getA, getB)
}