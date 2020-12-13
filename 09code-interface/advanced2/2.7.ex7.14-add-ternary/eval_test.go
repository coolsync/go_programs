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
		want string // 期望的结果
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
			// "172",	// 故意出错
		},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//!+Eval

		// additional ternary
		{"x ? y : z", Env{"x": 1, "y": 2, "z": 3}, "2"},
		{"x ? y : z", Env{"x": 0, "y": 2, "z": 3}, "3"},
	}

	var prevExpr string

	for _, d := range data {
		// Print expr only when it changes.
		if d.expr != prevExpr {
			fmt.Printf("\n%s:\n", d.expr) // 打印表达式
			prevExpr = d.expr
		}

		expr, err := Parse(d.expr) // 使用Parse包解析表示
		if err != nil {
			t.Error(err)
		}

		// get test result
		got := fmt.Sprintf("%.6g", expr.Eval(d.env)) // conv float result to string

		fmt.Printf("\t%v => %s\n", d.env, got)

		if got != d.want {
			t.Errorf("%s Eval(): %v = %s, want %s\n", d.expr, d.env, got, d.want)
		}
	}
}

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},

		// test ternary
		{"x ? y", "unexpected end of file"},
		{"x : y", "unexpected ':'"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}
