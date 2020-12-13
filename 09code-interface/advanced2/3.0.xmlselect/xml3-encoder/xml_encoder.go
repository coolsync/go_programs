package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func main() {
	type Address struct {
		City, State string
	}

	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	p := Person{Id: 1, FirstName: "bob", LastName: "kai", Age: 12}
	p.Comment = " Need more details. "
	p.Address = Address{"guangzhou", "north"}

	// create open file, 接受 编码后的文件
	f, err := os.OpenFile("./info.xml", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("open file err: ", err)
	}
	defer f.Close()

	// 创建编码器, 将go语言 data format conv xml format, 并写入到指定的输出流
	// enc := xml.NewEncoder(os.Stdout)

	enc := xml.NewEncoder(f)

	enc.Indent("", "\t")

	if err := enc.Encode(p); err != nil { // 编码操作
		log.Fatal(err)
	}

	fmt.Println("xml encoder: encode ok")
}
