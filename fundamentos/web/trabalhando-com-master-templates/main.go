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
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("web/trabalhando-com-master-templates/static/"))))
	router.HandleFunc("/{id}/view", ViewHandler)
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

func GetPostById(id string) Post {
	row := db.QueryRow("select * from posts where id=?", id)
	post := Post{}
	row.Scan(&post.Id, &post.Title, &post.Body)
	return post
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	view := template.Must(template.ParseFiles(
		"web/trabalhando-com-master-templates/templates/layout.html",
		"web/trabalhando-com-master-templates/templates/list.html",
	))
	if err := view.ExecuteTemplate(response, "layout.html", ListPosts()); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(response http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	view := template.Must(template.ParseFiles(
		"web/trabalhando-com-master-templates/templates/layout.html",
		"web/trabalhando-com-master-templates/templates/view.html",
	))
	if err := view.ExecuteTemplate(response, "layout.html", GetPostById(id)); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
