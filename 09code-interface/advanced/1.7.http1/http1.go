package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"Shoes": 50, "Socks": 10}
	log.Fatal(http.ListenAndServe(":8000", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "item: %s, price: %s\n", item, price)	// price格式化过程 自动调用 String method
	}
}
