package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

type IGoods interface {
	GetName() string
	GetPrice() float64
}

type Goods struct {
	Name  string
	Price float64
}

func (g *Goods) GetName() string {
	return g.Name
}

func (g *Goods) GetPrice() float64 {
	return g.Price
}

type Computer struct {
	// Name   string
	// Price  int
	Goods
	Cpu    string
	Memory int
	Disk   int
}

type Tshirt struct {
	// Name   string
	// Price  int
	Goods
	Gender uint8
	Color  string
	Size   int
}

type Car struct {
	// Name  string
	// Price int
	Goods
	Cap   int
	Power string
}

func EncodeObj2JsonFile(obj interface{}, filename string) bool {
	// 创建并打开文件
	dstFile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	// 创建dst file编码器
	encoder := json.NewEncoder(dstFile)

	if err := encoder.Encode(obj); err != nil {
		log.Fatal("encode failed: ", err)
		return false
	}

	log.Println("编码ok")
	return true
}

func main() {
	products := make([]interface{}, 0)
	// products := make([]IGoods, 0)

	products = append(products, Computer{Goods{"Rog", 10000}, "Ryzen", 16, 1024})
	products = append(products, Tshirt{Goods{"tshirt1", 1000}, 0, "red", 40})
	products = append(products, Car{Goods{"创冰", 100000}, 4, "全油"})

	for i, p := range products {
		// 一次性
		name := reflect.ValueOf(p).FieldByName("Name").Interface().(string)

		fmt.Printf("name: %d, %s\n", i, name)
		/* 
		// 通过反射 获取对象类型， （先前不知什么类型）
		objVal := reflect.ValueOf(p)
		// 获取 objVal 下 属性Name 
		nameVal := objVal.FieldByName("Name")
		// 获取 属性的值， 将反射变为正射（nil interface{}类型)
		nameInf := nameVal.Interface()
		// 断言string类型
		name := nameInf.(string) */

		// EncodeObj2JsonFile(p, "../../Files/"+name+".json")

		// EncodeObj2JsonFile(p, "../../Files/"+p.GetName()+".json")

		// name := p.(IGoods).GetName()
		// EncodeObj2JsonFile(p, "../../Files/"+name+".json")
	}

}
