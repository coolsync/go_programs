package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

func main() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	p := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`

	// 构造 xml 解码器, 将xml data format conv go lang data fromat, 读取指定的输入流
	dec := xml.NewDecoder(strings.NewReader(data))

	// 解码, 将data 写入 struct
	if err := dec.Decode(&p); err != nil {
		log.Fatal(err) // decode err
	}

	fmt.Printf("%#v\n\n", p)

	fmt.Printf("XMLName: %#v\n", p.XMLName)
	fmt.Printf("Name: %q\n", p.Name)
	fmt.Printf("Phone: %q\n", p.Phone)
	fmt.Printf("Email: %v\n", p.Email)
	fmt.Printf("Groups: %v\n", p.Groups)
	fmt.Printf("Address: %v\n", p.Address)
}
