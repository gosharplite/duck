package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	PORT  = flag.Int("port", 8080, "server port")
	DELAY = flag.Int("delay", 0, "delay in millisecond")
)

func main() {

	flag.Parse()

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *PORT), nil)

	fmt.Printf("ListenAndServe: %v\n", err)
}

func handler(w http.ResponseWriter, r *http.Request) {

	time.Sleep(time.Duration(*DELAY) * time.Millisecond)

	fmt.Fprint(w, "quack")
}
