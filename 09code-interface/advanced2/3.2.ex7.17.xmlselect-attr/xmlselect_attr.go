package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 创建 xml 解码器, 将 xml 解析	, search id=p3
	dec := xml.NewDecoder(os.Stdin)

	var stack []string
	var attrs []map[string]string

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			return
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push

			attr := make(map[string]string)

			for _, a := range tok.Attr {
				attr[a.Name.Local] = a.Value
			}

			attrs = append(attrs, attr) // 添加属性和对应值到 slice
			// fmt.Println(stack, attrs)

		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop, 丢弃尾部标签
			attrs = attrs[:len(attrs)-1]

		case xml.CharData: // 文本标签
			if containsAll(conv_slice(stack, attrs), os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}

		}
	}
}

func conv_slice(stack []string, attrs []map[string]string) []string {
	var result []string
	// for _, v := range stack {
	// 	result = append(result, v)
	// }

	// for i := range attrs {
	// 	for k, v := range attrs[i] {
	// 		result = append(result, k+"="+v)	// div class=d1
	// 	}
	// }

	for i, name := range stack {
		// 1
		result = append(result, name)
		// 2
		for attr, value := range attrs[i] {
			result = append(result, attr+"="+value)
		}

	}
	return result
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0] == y[0] {
			y = y[1:]	// 将切片原点向右移动一位, 丢掉 i 为 0 的元素, 继续和 x 元素一一比较
		}

		x = x[1:]
	}
	return false
}	
