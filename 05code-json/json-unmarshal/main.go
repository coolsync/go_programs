package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var jsonStr = `{"Name":"sam","Age":30,"Salary":3000.11,"Gender":true,"Bobbies":["吃东西","喝矿泉水","smokes"]}`

// json conv map
func json_conv_map() {

	m := make(map[string]interface{})

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		log.Fatal("map Unmarshal failed : ", err)
	}

	fmt.Printf("%v\n", m)
	// map[string]interface {}{"Age":30, "Bobbies":[]interface {}{"吃东西", "喝矿泉水", "smokes"}, "Gender":true, "Name":"sam", "Salary":3000.11}
}

// json conv struct
func json_conv_struct() {
	type Person struct {
		Name    string
		Age     int
		Salary  float64
		Gender  bool
		Bobbies []string
	}

	// p := Person{"sam", 30, 3000.11, true, []string{"吃东西", "喝矿泉水", "smokes"}}
	p := Person{}

	fmt.Printf("before unmarshal p address: %p\n", &p)
	if err := json.Unmarshal([]byte(jsonStr), &p); err != nil {
		log.Fatal("json Unmarshal struct failed : ", err)
	}

	fmt.Printf("after unmarshal p address: %p\n", &p)
	fmt.Println(p)
}

// json conv map slice
func json_conv_mapSli() {
	jsonData := `[{"gender":0,"hobbies":["看工程book","喝水","smoke 中华"],"name":"alice"},{"gender":1,"hobbies":["look women","喝牛奶","smoke 10￥"],"name":"bob"},{"gender":1,"hobbies":["吃东西","喝水笔","说 english"],"name":"tom"}]`

	ms := make([]map[string]interface{}, 3)

	if err := json.Unmarshal([]byte(jsonData), &ms); err != nil {
		log.Fatal("json Unmarshal struct failed : ", err)
	}

	fmt.Printf("ms: \n %+v\n", ms)
	fmt.Printf("ms: \n %#v\n", ms)


}

// json conv struct slice
func json_conv_structSli() {

	type Person struct {
		Name    string
		Age     int
		Salary  float64
		// Gender  bool
		Gender  int8
		Hobbies []string
	}

	jsonData := `[{"gender":0,"hobbies":["看工程book","喝水","smoke 中华"],"name":"alice"},{"gender":1,"hobbies":["look w","喝牛奶","smoke 10￥"],"name":"bob"},{"gender":1,"hobbies":["吃东西","喝水笔","说 english"],"name":"tom"}]`

	ps := make([]Person, 3)

	if err := json.Unmarshal([]byte(jsonData), &ps); err != nil {
		log.Fatal("json Unmarshal struct slice failed : ", err)
	}

	
	for i := range ps {
		fmt.Printf("ps %d : \n %+v\n",i, ps[i].Hobbies)
	}

	// get only attribute of name
	type Names struct {
		Name string
	}

	names := make([]Names, 3)

	if err := json.Unmarshal([]byte(jsonData), &names); err != nil {
		log.Fatal("json Unmarshal struct slice failed : ", err)
	}

	for i, v := range names {
		fmt.Printf("names %d : \n %#v\n",i, v)
	}
}

func main() {
	// json_conv_mapSli()
	json_conv_structSli()
}
