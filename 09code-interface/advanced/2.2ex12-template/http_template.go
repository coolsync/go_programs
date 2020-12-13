package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var db = database{"shoes": 20, "socks": 5}
var mux sync.Mutex

func main() {
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/price", db.price)


	log.Fatal(http.ListenAndServe(":8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (d database) list(w http.ResponseWriter, r *http.Request) {
	// parse tmpl
	t := template.Must(template.ParseFiles("index.html"))
	// write to tmpl
	if err := t.Execute(w, &d); err != nil {
		log.Fatalln(err)
	}
}

func (d database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if price, ok := d[item]; ok {
		fmt.Fprintf(w, "price: %s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %s isn't exist\n", item)
	}
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
