package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string
	Age int

	// embedded fields
	float32
	int
	next *Human
}

func main() {
	// 值包装struct
	rValue := reflect.ValueOf(Human{
		next: &Human{},
	})

	// get struct field nums
	fieldNum := rValue.NumField()
	fmt.Printf("field num is %v\n", fieldNum)

	// get index is 2 field
	field2 := rValue.Field(2)

	// judge index 2 field type
	fmt.Printf("field index 2 type is %v\n", field2.Type())

	// according to name get field
	fmt.Printf("%v\n", rValue.FieldByName("Age").Type())
	
	// according to rValue method 获取 next field 内 string field
	fmt.Printf("%v\n", rValue.FieldByIndex([]int{4, 0}).Type())
}