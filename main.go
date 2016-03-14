package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

var (
	PORT     = flag.Int("port", 80, "Server port")
	DELAY    = flag.Int("delay", 150, "delay in millisecond")
	LOGS     = flag.Bool("logs", false, "Show logs")
	LOCAL_IP = getLocalIP()
)

func main() {

	flag.Parse()

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *PORT), nil)

	fmt.Printf("ListenAndServe: %v\n", err)
}

func handler(w http.ResponseWriter, r *http.Request) {

	logs("start")

	time.Sleep(time.Duration(*DELAY) * time.Millisecond)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "quack - "+r.RemoteAddr+" -> "+LOCAL_IP+"\n")
	} else {
		fmt.Fprint(w, "quack - "+r.RemoteAddr+" -> "+LOCAL_IP+" : len(body)="+strconv.Itoa(len(body))+"\n")
	}

	logs("end")
}

func logs(format string, v ...interface{}) {
	if *LOGS {
		log.Printf(format, v...)
	}
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
