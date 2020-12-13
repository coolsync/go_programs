package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	data := []struct {
		expr string
		env  Env
		want string
	}{
		{
			"sqrt(A / pi)",
			Env{"A": 87616, "pi": math.Pi},
			"167",
		},
		{
			"pow(x, 3) + pow(y, 3)",
			Env{"x": 12, "y": 1},
			"1729",
		},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpl string

	for _, d := range data {
		if d.expr != prevExpl {
			// 无更改时, 打印表达式
			fmt.Printf("\n%s\n", d.expr)
			prevExpl = d.expr
		}

		// 解析 表达式
		expr, err := Parse(d.expr)
		if err != nil {
			t.Error(err) // parse err
		}
		got := fmt.Sprintf("%.6g", expr.Eval(d.env)) // 去掉小数点后的数字并返回

		fmt.Printf("\t%v => %s\n", d.env, got)

		if got != d.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", d.expr, d.env, got, d.want)
		}
	}

	/*
		for _, d := range data {
			expr, err := Parse(d.expr)
			if err == nil {
				err = expr.check(map[Var]bool{})
			}

			if err != nil {
				// 得到 err
				if err.Error() != d.want {
					t.Errorf("%s: got %q, want %q\n", d.expr, err, d.want)
				}
			}

			// got := fmt.Sprintf("%.6g", expr.Eval(d.env))

			// if got != d.want {
			// 	t.Errorf("%s: %v => %q, want %q\n", d.expr, d.env, got, d.want)
			// }
		} */
}
