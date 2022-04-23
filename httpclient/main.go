package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	apiUrl := "http://127.0.0.1:80/get"

	req, err := http.NewRequest("GET", apiUrl, nil)

	VERSION := os.Getenv("GOPATH")

	req.Header.Add("VERSION", VERSION)

	if err != nil {
		fmt.Println("Get error")
	}

	resp, _ := client.Do(req)

	for k, v := range resp.Header {
		fmt.Printf("%s=%s\n", k, v)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Get error")
	}

	fmt.Println(string(b))
}
