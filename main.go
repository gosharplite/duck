package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

var (
	PORT     = flag.Int("port", 80, "server port")
	DELAY    = flag.Int("delay", 0, "delay in millisecond")
	MAXCONNS = flag.Int("maxConns", 500, "max connections")
	LOCAL_IP = getLocalIP()
)

func main() {

	flag.Parse()

	http.Handle("/", NewLimitHandler(*MAXCONNS, handler))

	err := http.ListenAndServe(fmt.Sprintf(":%d", *PORT), nil)

	fmt.Printf("ListenAndServe: %v\n", err)
}

func handler(w http.ResponseWriter, r *http.Request) {

	t := time.Now()

	time.Sleep(time.Duration(*DELAY) * time.Millisecond)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "error: "+err.Error()+" , t="+time.Since(t).String()+"\n")
		return
	}

	fmt.Fprint(w, "quack - "+r.RemoteAddr+" -> "+LOCAL_IP+" , len(body)="+strconv.Itoa(len(body))+" , t="+time.Since(t).String()+"\n")
}

func getLocalIP() string {
	ip := "unknown"

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ip
	}

	for _, address := range addrs {

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}

	return ip
}

type limitHandler struct {
	connc   chan struct{}
	handler func(http.ResponseWriter, *http.Request)
}

func (h *limitHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	select {
	case <-h.connc:
		h.handler(w, req)
		h.connc <- struct{}{}
	default:
		//http.Error(w, "503 too busy", http.StatusServiceUnavailable)
	}
}

func NewLimitHandler(maxConns int, handler func(http.ResponseWriter, *http.Request)) http.Handler {
	h := &limitHandler{
		connc:   make(chan struct{}, maxConns),
		handler: handler,
	}
	for i := 0; i < maxConns; i++ {
		h.connc <- struct{}{}
	}
	return h
}
