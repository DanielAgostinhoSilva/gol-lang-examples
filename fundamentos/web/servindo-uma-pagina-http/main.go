package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello World")
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
