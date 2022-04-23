package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Info("Starting http server...")
	http.HandleFunc("/get", get)

	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	reqIp := r.Header.Get("ipv4")
	glog.Info(reqIp)

	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Add(k, v[0])
			fmt.Printf("%s=%s\n", k, v)
		}
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200 ok\n")
	glog.Info("200 ok")
}
