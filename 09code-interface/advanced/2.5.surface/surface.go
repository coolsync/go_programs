// The surface program plots the 3-D surface of a user-provided function.
package main

import (
	"fmt"
	eval "go-programs/09code-interface/advanced/2.4-eval/eval4"
	"io"
	"log"
	"math"
	"net/http"
)

// -- copied from gopl/ch3/surface --

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // x, y axis range (-xyrange...+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
)

var sin30, cos30 = 0.5, math.Sqrt(3.0 / 4.0) // sin(30°), cos(30°)

func corner(f func(x, y float64) float64, i, j int) (float64, float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) // compute surface height z

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(f, i+1, j)
			bx, by := corner(f, i, j)
			cx, cy := corner(f, i, j+1)
			dx, dy := corner(f, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

// 创建 parseAndCheck, 返回表达式 Expr 接口, err info
func parseAndCheck(s string) (eval.Expr, error) {
	// 判断 s 是否为空string
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}

	// 解析表达式
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("parse expr err: %v", err)
	}

	// 创建空的 vals 集合, 存放 无 err 的 Val set
	vals := make(map[eval.Var]bool)

	// 遍历vals, 判断x, y, r 的环境变量 是否存在
	for v := range vals {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("x, y, r env not exists")
		}
	}

	return expr, nil
}

// 创建 plot, 类似http.HandleFunc 将svg 返回给 client
func plot(w http.ResponseWriter, r *http.Request) {
	// 解析 多个 url params
	r.ParseForm()
	// 获取 expr后的params, 并使用 parseAndCheck 解析
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest) // 404
		return
	}

	// 设置头部info
	w.Header().Set("Content-Type", "image/svg+xml")

	// 使用surface 将 plot图结果 返回给 client
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)	//  distance from 0, 0
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
