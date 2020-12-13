package main

import (
	"encoding/xml"
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
	p.Address = Address{"guangzhou", "north"}
	p.Comment = "no more infos"
	/*
		xml_data, err := xml.Marshal(p) // go 语言数据格式 转 xml format
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", string(xml_data))
	*/

	xml_bs, err := xml.MarshalIndent(p, " ", "\t")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(xml_bs)	// 写入 st out stream
}

/*
// An Attr represents an attribute in an XML element (Name=Value).
// Attr表示XML元素（名称=值）中的属性。
type Name struct {
    Space, Local string
}
*/
