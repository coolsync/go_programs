package eval

import (
	"fmt"
	"math"
)

// 计算一个包含变量的表达式, 需要一个environment变量 将 变量的名字 映射 对应的值
type Env map[Var]float64

// Expr是一个算术表达式
type Expr interface {
	// Eval 在环境env上返回一个Expr的值
	Eval(env Env) float64
}

// 标识一个变量, e.g., x
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// 表示一个常量, e.g., 3.141
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// 表示一个一元运算符 Expr, e.g., -x.
type unary struct {
	op rune	// '+', '-'
	x Expr
}

// 实现 Expr interface, 获取表达式的值
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("不支持 unary operator %q", u.op))
}

// 表示一个二元运算符 表达式, e.g., x+y.
type binary struct {
	op rune // '+', '-', '/', '*'
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("不支持 binary operator %q", b.op))
}

// call 表示一个function call 表达式, e.g. math.Sqrt
type call struct {
	fn string
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		// c.args[1]是 num string时, 使用Parse解析后, 自动调用literal方法
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))	
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("不支持 function call %q", c.fn))
}



