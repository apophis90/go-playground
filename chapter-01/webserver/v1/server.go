package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		mutex.Lock()
		count++
		mutex.Unlock()
		fmt.Fprintf(rw, "URL.Path = %q\n", req.URL.Path)
	})
	http.HandleFunc("/count", func(rw http.ResponseWriter, req *http.Request) {
		mutex.Lock()
		fmt.Fprintf(rw, "Count: %d\n", count)
		mutex.Unlock()
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
