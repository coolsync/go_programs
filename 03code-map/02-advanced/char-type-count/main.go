package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

// 读取文本 count 各种 unicode 码点类型的

type class string

// define const
const (
	letter  class = "letter"
	number  class = "number"
	graphic class = "graphic"
	space   class = "space"
	symbol  class = "symbol"
)

func main() {
	class_count := make(map[class]int, 5)

	file, err := os.Open("./a.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// 构建 读取器
	reader := bufio.NewReader(file)

	for {
		r, _, err := reader.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stdin, " reader.ReadRune() : %s", err)
			return
		}

		// IsGraphic reports whether the rune is defined as a Graphic by Unicode.
		//  Such characters include letters, marks, numbers, punctuation, symbols, and spaces,
		//  from categories L, M, N, P, S, Zs.
		//IsGraphic报告是否通过Unicode将符文定义为图形。
		//这些字符包括字母，标记，数字，标点符号，符号和空格，
		//来自类别L，M，N，P，S，Zs。
		
		// IsSymbol reports whether the rune is a symbolic character
		// IsSymbol报告符文是否为符号字符
		switch {
		case unicode.IsLetter(r):
			class_count[letter]++
		case unicode.IsNumber(r):
			class_count[number]++
		case unicode.IsGraphic(r):
			class_count[graphic]++
		case unicode.IsSpace(r):
			class_count[space]++
		case unicode.IsSymbol(r):
			class_count[symbol]++
		}
	}

	// itrerator class_count

	for class, count := range class_count {
		fmt.Printf("count: %s\t --> count: %d\n", class, count)
	}
}
