package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	fmt.Fprintln(w, "hello world")
	fmt.Fprintln(w, r.Host)
	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.Method)
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe("0.0.0.0:20000", nil)
}
