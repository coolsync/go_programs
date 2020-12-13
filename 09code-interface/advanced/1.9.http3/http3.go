package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}

	mux := http.NewServeMux() // 多路选择器

	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	// 效果一样
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe(":8001", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (d database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "item: %s, price: %s\n", item, price)
	}
}

func (d database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := d[item]
	if !ok {
		// w.WriteHeader(http.StatusNotFound)	// 404
		// fmt.Fprintf(w, "no such item %q\n", item)
		msg := fmt.Sprintf("no such item %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "price: %s\n", price)
}
