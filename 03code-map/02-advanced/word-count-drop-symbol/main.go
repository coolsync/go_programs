package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {

	word_count := make(map[string]int)

	file, err := os.Open("./a.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	input := bufio.NewScanner(file)

	input.Split(bufio.ScanWords)

	for input.Scan() {

		str := input.Text()
		// 自动 隐式 转换 utf8 字符长度
		for i, r := range str {
			if unicode.IsPunct(r) {
				// fmt.Println("unicode 内", string(r))
				str = str[:i]
				break
			}
		}
		word_count[str]++
		
		// i := 0

		// 显示 转换 utf8 字符长度
		// for; i < len(str); {
		// 	// fmt.Println(i)
		// 	r, size := utf8.DecodeRuneInString(str[i:])

		// 	if unicode.IsPunct(r){
		// 		fmt.Println("unicode 内", string(r))
		// 		str = str[:i]
		// 		break
		// 	}

		// 	i += size
		// }

	}

	// iterator word_count
	for word, count := range word_count {
		fmt.Printf("word: %s \t --> count: %d\n", word, count)
	}
}
