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

		post := Post{Id: 1, Title: "Unnamed Post", Body: "No content"}

		if title := request.FormValue("title"); title != "" {
			post.Title = title
		}

		view := template.Must(template.ParseFiles("web/passando-valores-para-os-templates/templates/index.html"))
		if err := view.ExecuteTemplate(response, "index.html", post); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
