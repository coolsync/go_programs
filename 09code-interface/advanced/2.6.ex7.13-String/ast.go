package eval

// Expr 表达式运算 interface
type Expr interface {
	// 获取表达式 环境变量的值
	Eval(env Env) float64

	// check err expression, adds it to a set
	Check(vars map[Var]bool) error

	// 按一定格式 format output
	String() string
}

// 表达式中的变量名 在环境变量 映射成 对应的值
type Env map[Var]float64

// 一个 表达式 环境变量 例如: x
type Var string

// 一个常量的值, "3.141"
type literal float64

// a unary expression
type unary struct {
	op rune
	x  Expr
}

// a binary expression
type binary struct {
	op   rune
	x, y Expr // 获取表达式 环境变量的值
}

// 表示 a function call exprssion, "sin"
type call struct {
	fn   string
	args []Expr
}
