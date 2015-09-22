package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	PORT = flag.Int("port", 80, "Server port")
)

func main() {

	flag.Parse()

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *PORT), nil)
	
	fmt.Printf("ListenAndServe: %v\n", err)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("hit")

	fmt.Fprint(w, "quack")
}