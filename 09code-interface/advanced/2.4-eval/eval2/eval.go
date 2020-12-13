package eval

import (
	"fmt"
	"math"
)

// 将变量名Var(表达式) 映射成 对应值
type Env map[Var]float64

// 计算表达式的接口
type Expr interface {
	// 获取 此env环境内的 expl值
	Eval(env Env) float64

	// 在获取表达式的值之前, check expl值的正确性
	check(map[Var]bool) error
}

// Var 标识一个表达式 e.g., x
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v] // 返回对应值
}

// literal 是一个常量 表达式 e.g., 3.141
type literal float64

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

// 表示一个 一元 expresion
type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return u.x.Eval(env)
	case '-':
		return u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator %q", u.op))

}

// 表示一个 二元 expresion
type binary struct {
	op   rune
	x, y Expr
}

// get binary expl value
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env) // 可递归查询
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator %q", b.op))
}

// 表示一个 function call expl, e.g., sqrl
type call struct {
	fn   string
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call %q", c.fn))
}
