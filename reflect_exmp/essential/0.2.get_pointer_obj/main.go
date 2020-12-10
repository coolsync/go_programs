package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string
	Age int
}

func main() {
	// define a struct type pointer variable
	h := &Human{Name:"paul", Age: 1000}

	typeOfh := reflect.TypeOf(h)

	// get struct instantiation（实例化） 反射类型对象
	fmt.Printf("name: '%v', kind: '%v'\n", typeOfh.Name(), typeOfh.Kind())
	
	// 取类型对象
	typeOfh_elem := typeOfh.Elem()

	// 显示 反射类型对象 名称， 种类
	// display reflect.Type obj: name, kind
	fmt.Printf("element name: '%v', kind: '%v'\n", typeOfh_elem.Name(), typeOfh_elem.Kind())
}