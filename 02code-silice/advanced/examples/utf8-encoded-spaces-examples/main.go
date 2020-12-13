package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Exercise 4.6: Write an in-place function that
// squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
// in a UTF-8-encoded []byte slice into a single ASCII space.

// 创建一个功能 原地将一个UTF-8编码的[]byte类型的slice中相邻的空格
// （参考unicode.IsSpace）替换成一个空格返回


//IsSpace reports whether the rune is a space character as defined by Unicode's White Space property;
//IsSpace报告符文是否为Unicode的White Space属性定义的空格字符；

//  in the Latin-1 space this is 
//在Latin-1空间中，这是 

// '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).

// Other definitions of spacing characters are set by category Z and property Pattern_White_Space.
//间隔字符的其他定义由类别Z和属性Pattern_White_Space设置。


func replace(b []byte) []byte {
	for i := 0; i < len(b); {
		first, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(second) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
			}
		}
		i += size
	}
	return b
}

func dropSpace(b []byte) []byte {

	for i := 0; i < len(b); {
		first, size := utf8.DecodeRune(b[i:])

		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(b[i+size:])

			if unicode.IsSpace(second) {
				
				copy(b[i:], b[i+size:])	// 所有元素向前移动一位
				b = b[:len(b)-size]	// 舍弃最后一个元素

			}
		}

		i += size
	}
	return b
}

func main() {
	b := []byte("哈  哈哈  哈哈 aaa")

	
	res := dropSpace(b)

	fmt.Printf("%s\n", res)

	// for i := 0; i <len(b); {
	// 	first, size := utf8.DecodeRune(b[i:])

	// 	if unicode.IsSpace(first) {
			
	// 	}
	// 	fmt.Printf("%d-->%c\n", i, first)

	// 	fmt.Println(size)
	// 	i += size 
	// }
	
	// for i, v := range b {
	// 	fmt.Printf("%d-->%c\n", i, v)
	// }


}
