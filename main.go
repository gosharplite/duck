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
	LOCAL_IP = getLocalIP()
)

func main() {

	flag.Parse()

	http.HandleFunc("/", handler)

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
