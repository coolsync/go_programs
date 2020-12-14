package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		Name string
		Pro int `json:"pro" id:"100"`	// tag
	}

	// 创建Person实例
	p := Person{Name: "paul", Pro: 1000}
	
	// get Person reflect.Type obj
	typeOfp := reflect.TypeOf(p)

	// 遍历结构体成员
	for i := 0; i < typeOfp.NumField(); i++ {
		// 获取每个结构体成员类型
		filedType := typeOfp.Field(i)

		// 显示反射类型对象的name, tag
		fmt.Printf("filed -> name: %v\ttag: '%v'\n", filedType.Name, filedType.Tag)
	}

	// 通过filed name, find filed type info
	if PersonType, ok := typeOfp.FieldByName("Pro"); ok {

		// 从tag内获取对应tag
		fmt.Printf("tag: %v, %v\n", PersonType.Tag.Get("json"), PersonType.Tag.Get("id"))
	}
}

/* 
resutl:

filed -> name: Name	tag: ''
filed -> name: Pro	tag: 'json:"pro" id:"100"'
tag: pro, 100
*/