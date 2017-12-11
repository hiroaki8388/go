package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "world")
}

func main() {
	var server = http.Server{
		Addr: "127.0.0.1:8080"}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
