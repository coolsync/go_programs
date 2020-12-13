package main

import (
	"fmt"
	"log"
	"math"
)

/*
	求玩具球的 体积
	radius 负数， 直接恐慌
	处理 panic
	radius 不在 取值范围内[5, 50], 温和返回 （result, error）
	无论是 panic 还是 返回， 都使用 自定义异常InvalidRadiusError 实例 来操作
*/

// 封装 自定义 错误处理
type InvalidRadiusError struct {
	// 非法 半径
	radius float64
	// 可接受最小半径
	MinRadius float64
	// 可接受最大半径
	MaxRadius float64
}

// 实现 Error interface
func (ie *InvalidRadiusError) Error() string {
	err := fmt.Sprintf("非法半径 %.2f, 半径范围：[%.2f, %.2f]\n", ie.radius, ie.MinRadius, ie.MaxRadius)
	return err
}

// 使用工厂方法， 统一处理 InvalidRadiusError 类似部分, 只专注 输入值
func NewInvalidRadiusError(radius float64) *InvalidRadiusError {
	ie := new(InvalidRadiusError)

	ie.MinRadius = 5

	ie.MaxRadius = 50

	ie.radius = radius

	return ie
}

func GetBollVolumn(radius float64) (volumn float64, err error) {
	// radius 负数， 直接恐慌
	if radius < 0 {
		// panic(&InvalidRadiusError{radius, 5, 50})
		panic(NewInvalidRadiusError(radius))
	}

	// radius 不在 取值范围内[5, 50], 温和返回 （result, error）
	if radius < 5 || radius > 50 {
		// err := &InvalidRadiusError{radius, 5, 50}
		err := NewInvalidRadiusError(radius)
		return 0, err
	}

	return (4 / 3.0) * math.Pi * math.Pow(radius, 3), nil
}

func main() {

	// 在 功能 处理前， 找出恐慌的原因，阻止整个程序崩溃
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic :", err)
		}
	}()

	volumn, err := GetBollVolumn(-0.5)
	if err != nil {
		log.Fatal("get volumn failed: ", err)
	}

	fmt.Println("small boll volumn: ", volumn)
}
