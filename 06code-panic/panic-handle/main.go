package main

import (
	"fmt"
	"math"
)

/*
	求玩具球的 体积
	radius 负数， 直接恐慌
	处理 panic

*/

func getBollVolumn(radius float64) float64 {
	// 引起恐慌
	if radius < 0 {
		panic("半径不能为负")
	}
	return (4 / 3.0) * math.Pi * math.Pow(radius, 3)
}

func main() {

	// 在 功能 处理前， 找出恐慌的原因，阻止整个程序崩溃
	defer func(){
		res := recover()
		fmt.Println("panic :", res)
		fmt.Println("送房送钱送美女2")

	}()


	v := getBollVolumn(-1)
	
	fmt.Println("送房送钱送美女")
	fmt.Println("small boll volumn: ", v)
}