package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		view := template.Must(template.ParseFiles("web/renderizando-templates/templates/index.html"))
		if err := view.ExecuteTemplate(response, "index.html", nil); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
