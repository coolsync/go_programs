package eval

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{
			"-1 + -x",
			"(-1 + -x)",
		},

		{"-1 - x", "(-1 - x)"},

		{
			"sqrt(A / pi)",
			"sqrt((A / pi))",
		},

		{
			"pow(x,3)+pow(x,3)",
			"(pow(x, 3) + pow(x, 3))",
		},

		{"5 / 9 * (F - 32)", "((5 / 9) * (F - 32))"},
	}

	for i, tt := range tests {
		// parse expr
		expr, err := Parse(tt.expr)
		if err != nil {
			t.Error(err)
		}

		// 获取格式化输出方式
		got := expr.String()

		if got != tt.want {
			t.Fatalf("%d. got %s, want %s\n", i, got, tt.want)
		}
	}
}
