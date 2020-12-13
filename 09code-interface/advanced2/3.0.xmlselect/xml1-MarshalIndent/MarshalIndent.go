package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

/* 
  <person id="13">
      <name>
          <first>John</first>
          <last>Doe</last>
      </name>
      <age>42</age>
      <Married>false</Married>
      <City>Hanga Roa</City>
      <State>Easter Island</State>
      <!-- Need more details. -->
  </person>
*/
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

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	// MarshalIndent works like Marshal, 
	// but each XML element begins on a new indented line that starts with prefix 
	// and is followed by one or more copies of indent according to the nesting depth.
	// MarshalIndent的作用类似于Marshal，
	// ，但每个XML元素均以新的缩进行开头，该行以前缀开头
	// ，然后根据嵌套深度添加一个或多个缩进副本。 
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}