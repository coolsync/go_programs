// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

// func (f Fahrenheit) String() string {
// 	return fmt.Sprintf("%g℉", f)
// }

// func (k Kelvin) String() string {
// 	return fmt.Sprintf("%gK", k)
// }

// C_Conv_F converts a Celsius temperature to Fahrenheit.
func C_Conv_F(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// F_Conv_C converts a Fahrenheit temperature to Celsius.
func F_Conv_C(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 273.15
// K_Conv_C Kelvin conv  Celsius
func K_Conv_C(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// celsiusFlag 实现 flag.Value 接口
type celsiusFlag struct {
	Celsius
}

// 实现 Set方法
func (cf *celsiusFlag) Set(s string) error {
	var unit string   // 单位
	var value float64 // 传入的值

	fmt.Sscanf(s, "%f%s", &value, &unit) // 在格式化输出时, 连接传入值和单位

	switch unit {
	case "C", "℃":
		cf.Celsius = Celsius(value) // 更新 celsiusFlag value值
		return nil
	case "F", "℉":
		cf.Celsius = F_Conv_C(Fahrenheit(value))
		return nil
	case "K":
		cf.Celsius = K_Conv_C(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temp %q", s)
}

// 完成flag类型的创建
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	cf := celsiusFlag{Celsius: value}

	flag.CommandLine.Var(&cf, name, usage) // 传入地址时, 自动调用Set(String)

	return &cf.Celsius // 此时, 已完成对 cf 上 Celsius操作
}
