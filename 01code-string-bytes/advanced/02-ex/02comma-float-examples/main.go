package main

import (
	"bytes"
	"fmt"
)

// 需求： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Println(comma(os.Args[i]))
	// }

	fmt.Println(comma("-1222343412.345432531253"))
}

func comma(str string) string {
	n := len(str)

	if n <= 3 {
		return str
	}
	
	var buf bytes.Buffer
	
	s := []byte(str)

	// 判断 前面的 正负号
	sign := s[0]

	if sign == '+' || sign == '-' {
		s = s[1:]
		buf.WriteByte(sign)
	}

	// 负数后面不用分割
	last := make([]byte, 0)

	for i, v := range s {
		if v == '.' {
			last = s[i:]
			s = s[:i]
		}
	}


	for i, v := range s {
		// 整数判断
		// if i%3 == 0 && i != 0 {
		// 	buf.WriteString(",")
		// }
		if i%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", v)
	}

	buf.Write(last)
	// buf.WriteString(string(last))
	return buf.String()
}

// func comma(s string) string {
// 	if s == "" {
// 		return ""
// 	}
// 	str := []byte(s)
// 	var buf bytes.Buffer
// 	sign := str[0]
// 	if sign == '+' || sign == '-' {
// 		buf.WriteByte(sign)
// 		str = str[1:]
// 	}

// 	last := make([]byte, 0)
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == '.' {
// 			last = str[i:]
// 			str = str[:i]
// 		}
// 	}

// 	for i := 0; i < len(str); i++ {
// 		if (len(str)-i)%3 == 0 {
// 			buf.WriteByte(',')
// 		}
// 		buf.WriteByte(str[i])
// 	}
// 	buf.WriteString(string(last))
// 	return buf.String()
// }
