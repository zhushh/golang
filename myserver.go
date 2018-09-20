package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	DEFAULT_IP   = "localhost"
	DEFAULT_PORT = "8000"
)

var mux sync.Mutex
var count int = 0

var IP string
var PORT string

func main() {
	IP = GetIpFromCmd()
	PORT = GetPortFromCmd()

	http.HandleFunc("/", handler_nice)
	http.HandleFunc("/count", handler_count)
	log.Fatal(http.ListenAndServe(IP+":"+PORT, nil))
}

func GetIpFromCmd() string {
	for i, v := range os.Args {
		if strings.HasPrefix(v, "-h") && v != "-h" {
			return v[2:]
		} else if v == "-h" {
			return os.Args[i+1]
		}
	}

	return DEFAULT_IP
}

func GetPortFromCmd() string {
	for i, v := range os.Args {
		if strings.HasPrefix(v, "-p") && v != "-p" {
			return v[2:]
		} else if v == "-p" {
			return os.Args[i+1]
		}
	}

	return DEFAULT_PORT
}

func handler_nice(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Receive request: %v\n", r)
	mux.Lock()
	count++
	mux.Unlock()
	fmt.Fprintf(w, "request path = %s\n", r.URL)
}

func handler_count(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Receive request: %v\n", r)
	fmt.Fprintf(w, "count = %d\n", count)
}
