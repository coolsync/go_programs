package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 20, "socks": 5}
	http.HandleFunc("/list", db.list)	// 传入 method value
	http.HandleFunc("/price", db.price)	

	log.Fatal(http.ListenAndServe(":8001", nil))	// 为nil，将使用DefaultServeMux 作为 默认处理方式
}

// 将 price 格式化输出
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// 数据库 物品 映射 价格
type database map[string]dollars

// 获取物件与价格list
func (d database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "item %s, price %s\n", item, price)
	}
}

// 获取指定 item 的 price
func (d database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	price, ok := d[item]
	if !ok {
		msg := fmt.Sprintf("no such item %q\n", item)
		http.Error(w, msg, http.StatusNotFound)	// 404
		return
	}

	fmt.Fprintf(w, "price: %s\n", price)
}
