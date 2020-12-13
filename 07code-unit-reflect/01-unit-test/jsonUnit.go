package mathUnit

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Salary  float64
	Gender  uint8
	Hobbies []string
}

// 编码
func EncodePerson2JsonFile(filename string, p *Person) bool {

	// 创建并打开文件
	dstFile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer dstFile.Close()

	// 构建目标文件编码器
	encoder := json.NewEncoder(dstFile)

	if err := encoder.Encode(p); err != nil {
		fmt.Println("编码 failed: ", err)
		return false

	}

	// fmt.Println("EncodePerson2JsonFile 编码成功")
	return true
}

// 解码器
func DecodeJsonFile2Person(filename string) (*Person, error) {

	srcFile, _ := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer srcFile.Close()
	// p := Person{}
	pPtr := new(Person)

	// 构建 源文件解码器
	decoder := json.NewDecoder(srcFile)

	if err := decoder.Decode(pPtr); err != nil {
		// fmt.Println("解码 failed: ", err)
		return nil, errors.New("解码 failed")
	}

	// fmt.Println("DecodeJsonFile2Person 解码成功")

	return pPtr, nil
}
