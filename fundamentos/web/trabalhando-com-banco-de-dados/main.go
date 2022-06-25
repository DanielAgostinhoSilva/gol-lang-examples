package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"text/template"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:root@/go_course?charset=utf8")

func main() {

	stmt, err := db.Prepare("Insert into posts(title, body) values (?, ?)")
	checkErr(err)

	_, err = stmt.Exec("My First Post", "My First content")
	checkErr(err)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		post := Post{Id: 1, Title: "Unnamed Post", Body: "No content"}

		if title := request.FormValue("title"); title != "" {
			post.Title = title
		}

		view := template.Must(template.ParseFiles("web/trabalhando-com-banco-de-dados/templates/index.html"))
		if err := view.ExecuteTemplate(response, "index.html", post); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}