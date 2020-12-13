// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}

// C_Conv_F converts a Celsius temperature to Fahrenheit.
func C_Conv_F(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// F_Conv_C converts a Fahrenheit temperature to Celsius.
func F_Conv_C(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

/*
package flag
// Value is the interface to the value stored(储存) in a flag.

type Value interface {
	String() string
	Set(string) error
}
*/

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string   // value后的单位
	var value float64 // 传入的具体值

	// Sscanf scans the argument string, storing successive space-separated
	fmt.Sscanf(s, "%f%s", &value, &unit) // 将值和单位连接一个字符串

	switch unit {
	case "C", "℃":
		f.Celsius = Celsius(value)
		return nil
	case "F", "℉":
		f.Celsius = F_Conv_C(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invaid temp %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}

	flag.CommandLine.Var(&f, name, usage) // 传入f的地址, 自动调用Set方法

	return &f.Celsius
}
