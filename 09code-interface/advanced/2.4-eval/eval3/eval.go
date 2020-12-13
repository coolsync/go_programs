package eval

import (
	"fmt"
	"math"
	"strings"
)

// create implement expression Expr interface various types(各种类型)

// 算术表达式
type Expr interface {
	// return 环境变量 Expr 对应的值
	Eval(env Env) float64

	// Check error 在 Expr内, 将它的值Vars添加到集合
	Check(vals map[Var]bool) error
}

// 计算一个包含变量的表达式, 需要一个environment变量 将 变量的名字 映射 对应的值
type Env map[Var]float64

// 标识一个变量名, e.g., x
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

// 表示一个常量, e.g., 3.141
type literal float64

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

// 表示一个一元表达式, e.g., '+'
type unary struct {
	op rune
	x  Expr
}

// 获取 一元表达式 对应的值
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}

	panic(fmt.Sprintf("unsorpport unary operator %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected %q", u.op)
	}
	return u.x.Check(vars)
}

// 表示一个二元表达式, e.g., '+ - / *'
type binary struct {
	op   rune
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
	panic(fmt.Sprintf("unsorpport binary operator %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected %q", b.op)
	}

	if err := b.x.Check(vars); err != nil {
		return err
	}

	return nil
}

// 表示一个 function call expresstion	'pow'
type call struct {
	fn   string // function name
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
	panic(fmt.Sprintf("unsorpport function call %q", c.fn))
}

func (c call) Check(vars map[Var]bool) error {
	// 1. judge c.fn 是否存在, 返回 func name 的 params num
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	// 2. judge params 是否一致
	if arity != len(c.args) {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}
	// 3. 对每一个arg, check error
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
