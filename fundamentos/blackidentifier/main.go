package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://google.com")
	body, _ := ioutil.ReadAll(res.Body)
	err := res.Body.Close()
	if err != nil {
		return
	}
	fmt.Printf("%s", body)
}
