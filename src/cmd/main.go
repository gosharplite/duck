package main

import (
	"flag"
	"fmt"
	pro "instrument"
	"log"
	"net/http"
	"time"
)

var (
	PORT  = flag.Int("port", 8080, "server port")
	DELAY = flag.Int("delay", 0, "delay in millisecond")
)

func main() {

	flag.Parse()

	go func() {
		h := pro.NewInstrument()

		err := h.Serve()
		if err != nil {
			log.Fatalf("instrumentation failed to start: %v", err)
		}
	}()

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *PORT), nil)
	if err != nil {
		log.Fatalf("duck server failed to start: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	timer := pro.NewTimer("handler")
	defer timer.ObserveDuration()

	time.Sleep(time.Duration(*DELAY) * time.Millisecond)

	fmt.Fprint(w, "quack")
}
