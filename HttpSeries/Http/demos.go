package Http
// package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func HttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}


func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}

