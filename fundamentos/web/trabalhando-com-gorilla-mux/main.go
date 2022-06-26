package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	//stmt, err := db.Prepare("Insert into posts(title, body) values (?, ?)")
	//checkErr(err)
	//
	//_, err = stmt.Exec("My First Post", "My First content")
	//checkErr(err)

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	fmt.Println(http.ListenAndServe(":8080", router))
}

func ListPosts() []Post {
	rows, err := db.Query("Select * from posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}

	return items
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	view := template.Must(template.ParseFiles("web/trabalhando-com-gorilla-mux/templates/index.html"))
	if err := view.ExecuteTemplate(response, "index.html", ListPosts()); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
