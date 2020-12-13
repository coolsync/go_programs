package main

import (
	"bufio"
	"fmt"
	"go-programs/09code-interface/advanced2/2.8.ex7.15-reader-calc/eval"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("input expr ")
	r := bufio.NewReaderSize(os.Stdin, 4096)

	// 1. handle expression
	for {
		fmt.Print("> ")
		line, err := r.ReadBytes('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatal(err)
		}

		expr, err := eval.Parse(string(line))
		if err != nil {
			fmt.Printf("invalid expression %s\n", line)
			continue
		}
		/*
			fmt.Printf("expr: %#v\n", expr)
			pow(x,y) {pow [x y]} expr: eval.call{fn:"pow", args:[]eval.Expr{"x", "y"}}
		*/

		// use Check method, 将没有error expr 环境变量名 添加 set
		vars := map[eval.Var]bool{}

		// 2. get expr env val name
		err = expr.Check(vars)
		if err != nil {
			fmt.Printf("invalid expression %s\n", line)
			continue
		}

		/*
			fmt.Printf("vals: %#v\n", vars)
			> pow(x,y)
			vals: map[eval.Var]bool{"x":true, "y":true}
			> x - y
			vals: map[eval.Var]bool{"x":true, "y":true}
		*/

		// 3. handle env 对应值
		env := eval.Env{}

		for v := range vars {

			for {
				fmt.Printf("%s: ", v)
				// line, err := r.ReadBytes('\n')
				line, _, err := r.ReadLine()
				if err == io.EOF {
					return
				} else if err != nil {
					log.Fatal(err)
				}

				f, err := strconv.ParseFloat(string(line), 64)
	
				if err != nil {
					fmt.Printf("invalid value %s\n", line)
					continue // re input
				}

				env[v] = f
				break // get val after, break for
			}
		}

		fmt.Printf("expr result: %v\n", expr.Eval(env))
	}
}
