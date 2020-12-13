package main

import (
	"go-programs/09code-interface/advanced2/2.9.ex7.16-web-calc/eval"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	http.HandleFunc("/", index) // handle request path
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe(":8000", nil)) // 绑定 listen port
}

func index(w http.ResponseWriter, req *http.Request) {
	// define and parse html tmpl
	t := template.Must(template.ParseFiles("./index.html"))

	// render html tmpl
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

// get post req data
func calc(w http.ResponseWriter, req *http.Request) {
	// define and parse html tmpl
	t := template.Must(template.ParseFiles("./index.html"))

	expr_str := req.PostFormValue("expr") // get client set post data
	env_str := req.PostFormValue("env")

	// parse expression
	expr, err := eval.Parse(expr_str)
	if err != nil {
		log.Fatal(err)
	}

	// parse expr env value
	env, err := parseEnv(env_str)
	if err != nil {
		log.Fatal("parseEnv: ", err)
	}

	// render html tmpl
	if err := t.Execute(w, expr.Eval(env)); err != nil {
		log.Fatal("t.Execute: ", err)
	}
}

func parseEnv(env_str string) (eval.Env, error) {
	env := make(map[eval.Var]float64)

	// clean 多余 字符
	fields := strings.FieldsFunc(env_str, func(r rune) bool {
		return strings.ContainsRune(`{}:=,\"`, r) || unicode.IsSpace(r)
	})

	// 提取 剩余各个对应字符
	for i := 0; i+1 < len(fields); i += 2 {
		k := strings.TrimSpace(fields[i])
		v := strings.TrimSpace(fields[i+1])

		// 将 k 对应 v 转为 float
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}

		env[eval.Var(k)] = val
	}

	return env, nil
}
