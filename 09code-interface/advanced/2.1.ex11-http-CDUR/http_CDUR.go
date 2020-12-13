package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mux sync.Mutex

func main() {
	db := database{"shoes": 30, "socks": 5}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/read", db.read)

	// 绑定端口
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// 格式化输出 $
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// get item, price list
func (d database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "item: %s, price: %s\n", item, price)
	}
}

// get price
func (d database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := d[item] // check item 是否 存在
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%q isn't exist\n", item)
		return
	}
	fmt.Fprintf(w, "item: %s, price: %s\n", item, price)
}

// create item, price
func (d database) create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("param err: %v\n", err)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}

	mux.Lock()
	_, ok := d[item]
	if ok {
		fmt.Fprintf(w, "%q is exist!\n", item)
	} else {
		d[item] = dollars(p)
	}

	mux.Unlock()
}

// delele item
func (d database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	_, ok := d[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%q isn't exist\n", item)
	} else {
		mux.Lock()
		delete(d, item)
		mux.Unlock()
	}
}

// update item price
func (d database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "param err: %v\n", err)
		return
	}

	mux.Lock()

	_, ok := d[item]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %s isn't exist\n", item)
	} else {
		d[item] = dollars(p)
	}

	mux.Unlock()
}

// read item
func (d database) read(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	price, ok := d[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %s isn't exist\n", item)
		return
	}

	fmt.Fprintf(w, "item: %s, price: %s\n", item, price)
}
