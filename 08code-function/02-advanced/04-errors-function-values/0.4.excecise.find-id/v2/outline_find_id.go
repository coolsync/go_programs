package main

import (
	"flag"
	"fmt"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var id = flag.String("id", "", "请输入id值")

var f func(i int) int

func main() {
	
	fmt.Println(f == nil)

	f = demo
	fmt.Println(f==nil)

	flag.Parse()
	fmt.Println("*id: ", *id)
	fmt.Println(flag.Args())
	for _, url := range flag.Args() {
		fmt.Println("url: ", url)
		fmt.Println("*id: ", *id)
	}

}

func demo(i int) int {
	fmt.Println("djfls")
	return i
}

// func relace(s string) string {
// 	return s
// }
