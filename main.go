package main

import(
	"fmt"
	"net/http"
)

var db *DB

func main() {
	db = NewDB()

	http.HandleFunc("/", Hello)
	http.ListenAndServe(":5001", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
}