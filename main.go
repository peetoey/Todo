package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID          int
	Subject     string
	Description string
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {

	http.HandleFunc("/", todo)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}

func todo(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "postgres://zecolkej:KwZAmgZmsXEN_xPhnzWEsNmNRz2rNlJl@john.db.elephantsql.com:5432/zecolkej")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, subject, description FROM todo")
	if err != nil {
		log.Fatal(err)
	}

	resp := Todo{}
	var results []Todo

	for rows.Next() {
		var id int
		var sj, des string
		err = rows.Scan(&id, &sj, &des)
		resp.ID = id
		resp.Subject = sj
		resp.Description = des

		results = append(results, resp)

		if err != nil {
			log.Fatal(err)
		}
	}

	templates.Execute(res, results)
}

func delete(res http.ResponseWriter, req *http.Request) {
	fmt.Println("begin delete")
	db, err := sql.Open("postgres", "postgres://zecolkej:KwZAmgZmsXEN_xPhnzWEsNmNRz2rNlJl@john.db.elephantsql.com:5432/zecolkej")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	key := req.URL.Query().Get("ID")

	_, err = db.Exec(fmt.Sprintf("DELETE FROM todo WHERE id=%s", key))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted")

	http.Redirect(res, req, "/", 301)
}
