package eval

import (
	"fmt"
	"strings"
)

func (v Var) check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (l literal) check(vars map[Var]bool) error {
	return nil
}

// 检查 unary operator 
func (u unary) check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary operator %q", u.op)
	}

	if err := u.x.check(vars); err != nil {
		return err
	}

	return nil
}

func (b binary) check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected unary operator %q", b.op)
	}

	if err := b.x.check(vars); err != nil {
		return err
	}

	return b.y.check(vars)
}

// 创建 map, 指定 func name and func args
var paramsNum = map[string]int{"pow": 2, "sin": 1, "sqrt":1}

// check function call expl
func (c call) check(vars map[Var]bool) error {
	arity, ok := paramsNum[c.fn]
	if !ok {
		return fmt.Errorf("unexpected function call %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("func call args num %d, want %d", arity, len(c.args))
	}

	for _, arg := range c.args {
		if err := arg.check(vars); err != nil {
			return err
		}
	}
	return nil
}