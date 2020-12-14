package main

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	zero Enum = 0
)

type Human struct {
	Name string
	Age  int
}

func main() {
	var h Human

	// get struct reflect.Type obj
	typeOfh := reflect.TypeOf(h)

	// 获取反射对象的名称， 种类
	fmt.Println(typeOfh.Name(), typeOfh.Kind())

	// get zero reflect.Type obj
	typeOfZero := reflect.TypeOf(zero)

	// 获取反射对象的名称， 种类
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind())
}
