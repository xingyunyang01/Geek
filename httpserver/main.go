package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	flag.Parse()
	log.Println("Starting http server...")
	http.HandleFunc("/get", get)

	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	reqIp, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))

	if err == nil {
		log.Println(reqIp)
	}

	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Add(k, v[0])
			fmt.Printf("%s=%s\n", k, v)
		}
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200 ok\n")
	log.Println("200 ok")
}
