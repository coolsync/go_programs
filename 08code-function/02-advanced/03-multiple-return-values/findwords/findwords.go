package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

// 1. get file name in os.Args
// 2. create findwords func, passed fname return map words
// 3. open  file
// 4. use newScanner get scanner, for add word to map

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "arg have to 1")
		os.Exit(1)
	}

	words := make([]string, 0)

	for _, fname := range os.Args[1:] {

		m := findwords(fname)
		for k := range m {
			// fmt.Printf("word: %s --> nums: %d\n", k, v)
			words = append(words, k)
		}
		sort.Strings(words)

		for _, v := range words {
			fmt.Printf("word: %s --> nums: %d\n", v, m[v])
		}
	}
}

func findwords(fname string) map[string]int {
	wordsCount := make(map[string]int)

	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findwords: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	input := bufio.NewScanner(file)

	input.Split(bufio.ScanWords)

	for input.Scan() {
		str := input.Text()

		// 自动 隐式 转换 utf-8 字符
		// IsPunct judge symbol, unicode.IsSpace(v) judge space
		for i, r := range str {
			if unicode.IsPunct(r) {
				str = str[:i]
				break
			}
		}
		// fmt.Println(str)
		wordsCount[str]++
	}

	return wordsCount
}
