package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// encoder map format write to json file
func map_to_jsonFile() {

	// sam 的 go 格式 转为 json 格式
	m := make(map[string]interface{})
	m["name"] = "sam"
	m["age"] = 30
	m["salary"] = 2000.00
	m["gender"] = true
	m["hobbies"] = []string{"吃东西", "喝矿泉水", "smokes"}

	// create and open file
	dstFile, err := os.OpenFile("../Files/db.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatal("open file failed err: ", err)
	}

	defer dstFile.Close()

	// 创建 编码器， 将编码后的 map 数据（json data） 写入 josn文件
	enc := json.NewEncoder(dstFile)

	// encdoe 将 map 数据  转换 json 数据
	if err := enc.Encode(&m); err != nil {
		log.Fatal("encode failed err: ", err)
	}
	
	fmt.Println("编码器编码 ok")
}


// encoder struct slice format write to json file
func structSli_to_jsonFile() {
	// 创建并实例化结构体
	type Person struct {
		Name    string
		Age     int
		Salary  float64
		Gender  uint8
		Hobbies []string
	}

	p1 := Person{"alice", 20, 3000.00, 0, []string{"看工程book", "喝水", "smoke 中华"}}
	p2 := Person{"bob", 30, 2000.00, 1, []string{"看女的", "喝牛奶", "smoke 10￥"}}
	p3 := Person{"tom", 31, 2500.00, 1, []string{"吃东西", "喝水笔", "说 english"}}


	ps := make([]Person, 0)
	ps = append(ps, p1, p2, p3)

	// 创建并打开文件， 将编码后的json数据写入目标文件
	dstFile, err := os.OpenFile("../Files/db2.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("create and open file failed err: ", err)
	}

	// 构建编码器， 将编码后的json数据 写入到指定的写入流
	encoder := json.NewEncoder(dstFile)

	// 编码器 编码 struct slice data, ok后 转为 json data 格式
	if err := encoder.Encode(&ps); err != nil {
		log.Fatal("编码 failed err: ", err)
	}

	fmt.Println("编码器编码 ok")
}


// decoder json file to map 
func jsonFile_to_map() {

	// 打开json数据的源文件
	srcFile, err := os.OpenFile("../Files/db.json", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("open json srcFile failed err: ", err)
	}
	defer srcFile.Close()

	// 创建map 接受解码器解码的数据
	dataMap := make(map[string]interface{})

	// 构建 源文件解码器
	decoder := json.NewDecoder(srcFile)

	// 解码 json文件的数据， 丢入到 存储 dataMap 的内存上
	if err := decoder.Decode(&dataMap); err != nil {
		log.Fatal("解码 failed err: ", err)
	}

	fmt.Println("解码 ok: ", dataMap)

}

// decoder decode json file to struct slice
func jsonFile_to_structSli() {
	// 打开存储 json data 的 源文件
	srcFile, err := os.OpenFile("../Files/db2.json", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("open json srcFile failed err: ", err)
	}
	defer srcFile.Close()

	// 创建用于接收 数据 结构体
	type Person struct {
		Name    string
		Age     int
		Salary  float64
		Gender  uint8
		Hobbies []string
	}
	
	// 创建 struct slice， 接受 解码数据
	ps := make([]Person, 0)

	// 构建 源文件解码器
	decoder := json.NewDecoder(srcFile)

	// 解码 json数据后，丢入 存储 ps 的内存上
	if err := decoder.Decode(&ps); err != nil {
		log.Fatal("解码 failed err: ", err)

	}

	// 返回解码后的数据
	fmt.Println("解码 ok: \n", ps)

	for i, v := range ps {
		fmt.Printf("\nps slice %d:\n %+v\n", i, v)
	}
}


func main() {
	// encoder struct slice format write to json file
	// structSli_to_jsonFile()

	// decoder decode json file to map 
	// jsonFile_to_map()

	// decoder decode json file to struct slice
	jsonFile_to_structSli()
}
