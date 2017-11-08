package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "%s %s %s\n", req.Method, req.URL, req.Proto)
		for key, val := range req.Header {
			fmt.Fprintf(rw, "Header[%q] = %q\n", key, val)
		}
		fmt.Fprintf(rw, "Host = %q\n", req.Host)
		fmt.Fprintf(rw, "RemoteAddr = %q\n", req.RemoteAddr)
		if err := req.ParseForm(); err != nil {
			log.Print(err)
		}
		for key, val := range req.Form {
			fmt.Fprintf(rw, "Form[%q] = %q\n", key, val)
		}
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
