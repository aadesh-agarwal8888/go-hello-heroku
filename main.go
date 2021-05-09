package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello WOrld"))
}
