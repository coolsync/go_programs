package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 20, "socks": 5}
	log.Fatal(http.ListenAndServe(":8001", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path { // 获取path
	case "/list":
		for item, price := range d {
			fmt.Fprintf(w, "item: %s, price: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item") // get param
		price, ok := d[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item %q", item)
			return
		}
		fmt.Fprintf(w, "item: %s, price: %s", item, price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such link %q", r.URL)
	}
}
