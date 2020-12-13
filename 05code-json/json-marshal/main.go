package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// create Person strcut
// 必须大写 才能导出到 json format
type Person struct {
	Name    string
	Age     int
	Salary  float64
	Gender  bool
	Bobbies []string
}

// use struct conv json
func struct_conv_json() {

	p := Person{"sam", 30, 3000.11, true, []string{"吃东西", "喝矿泉水", "smokes"}}

	data, err := json.Marshal(p)

	if err != nil {
		log.Fatal("struct conv json 失败: ", err)
	}

	fmt.Println("json marshal output:",string(data))

}

// use map conv josn
func map_conv_json() {
	m := make(map[string]interface{})

	m["name"] = "sam"
	m["age"] = 30
	m["salary"] = 2000.00
	m["gender"] = true
	m["hobbies"] = []string{"吃东西", "喝矿泉水", "smokes"}

	data, err := json.Marshal(m)

	if err != nil {
		log.Fatal("map conv json 失败: ", err)
	}

	fmt.Println(string(data))
}

// use map slice conv josn, description sam colleague
func mapSli_conv_json() {
	dataSli := make([]map[string]interface{}, 0)

	m := make(map[string]interface{})
	m["name"] = "alice"
	m["gender"] = 0
	m["hobbies"] = []string{"看工程book", "喝水", "smoke 中华"}

	m2 := make(map[string]interface{})
	m2["name"] = "bob"
	m2["gender"] = 1
	m2["hobbies"] = []string{"看女的", "喝牛奶", "smoke 10￥"}

	m3 := make(map[string]interface{})
	m3["name"] = "tom"
	m3["gender"] = 1
	m3["hobbies"] = []string{"吃东西", "喝水笔", "说 english"}

	dataSli = append(dataSli, m, m2, m3)

	data, err := json.Marshal(dataSli)

	if err != nil {
		log.Fatal("map conv json 失败: ", err)
	}

	fmt.Println(string(data))
}


func main() {
	mapSli_conv_json()
}
