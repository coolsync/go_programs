// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

// Celsius conv Fahrenheit
func C_Conv_F(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Fahrenheit conv Celsiusa
func F_Conv_C(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

func main() {

	fmt.Printf("摄氏 沸水温度值: %g\n", BoilingC-FreezingC) // 100

	boilingF := C_Conv_F(BoilingC) // 华氏 沸水温度

	fmt.Printf("华氏 沸水温度值: %g\n", boilingF-C_Conv_F(FreezingC)) // 180

	// fmt.Printf("华氏 沸水温度值: %g\n", boilingF-FreezingC)
	// invalid operation: boilingF - FreezingC (mismatched types Fahrenheit and Celsius)

	fmt.Println("----------------")

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"

	// fmt.Println(c == f)
	// invalid operation: c == f (mismatched types Celsius and Fahrenheit)

	fmt.Println(c == Celsius(f)) // "true"!

	fmt.Println("----------------")

	c = F_Conv_C(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c) // 隐式自动调用String()方法
	fmt.Printf("%s\n", c)
	fmt.Println(c)

	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}
