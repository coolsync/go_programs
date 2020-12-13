// ast pkg define interface and struct
package eval

// Expr is expression 运算
type Expr interface {
	// get expression 环境变量名对应的值
	Eval(env Env) float64
	// resports exprssion 出现 error, adds it to set
	Check(vars map[Var]bool) error
}

// Var is expression 环境变量名
type Var string

// literal is expression 常量值, eg. 3.141
type literal float64

// unary 一元运算符
type unary struct {
	op rune
	x Expr
}
// binary 二元运算符
type binary struct {
	op rune
	x,y Expr
}

// function call eg. "sin"
type call struct {
	fn string
	args []Expr
}