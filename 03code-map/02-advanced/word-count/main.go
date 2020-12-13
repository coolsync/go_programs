package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 统计 文本 单个单词 出现的 次数

func main() {
	word_count := make(map[string]int)

	file, err := os.Open("./a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := bufio.NewScanner(file)

	input.Split(bufio.ScanWords)

	str := ""
	for input.Scan() {
		str = strings.ReplaceAll(input.Text(), ",", "")
		str = strings.ReplaceAll(str, "。", "")
		word_count[str]++
	}

	for word, count := range word_count {
		fmt.Printf("word: %s\tcount: %d\n", word, count)
	}
}
