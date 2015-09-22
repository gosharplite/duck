package main

import (
	"time"
	"flag"
	"fmt"
	"net/http"
	"log"
)

var (
	PORT = flag.Int("port", 80, "Server port")
	DELAY = flag.Int("delay", 150, "delay in millisecond")
)

func main() {

	flag.Parse()

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *PORT), nil)
	
	fmt.Printf("ListenAndServe: %v\n", err)
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Printf("start")
	
	time.Sleep(time.Duration(*DELAY) * time.Millisecond)
	
	log.Printf("end")

	fmt.Fprint(w, "quack")
}