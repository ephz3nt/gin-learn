package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("hello.txt")
	fmt.Fprintln(w, string(f))
}

func main() {
	addr := "0.0.0.0:20000"
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(addr, nil)
	fmt.Println("server start on: ", addr)
	if err != nil {
		fmt.Println("server start failed, err: ", err)
		return
	}

}
