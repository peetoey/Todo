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

var (
	db        *sql.DB
	err       error
	templates = template.Must(template.ParseFiles("index.html"))
)

func main() {
	err = dbConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	http.HandleFunc("/", todo)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}

func todo(res http.ResponseWriter, req *http.Request) {

	// Query fro db
	rows, err := db.Query("SELECT id, subject, description FROM todo")
	if err != nil {
		log.Fatal(err)
	}

	todo := Todo{}
	var resp []Todo

	// Read data increase to array
	for rows.Next() {
		var id int
		var sj, des string
		err = rows.Scan(&id, &sj, &des)
		todo.ID = id
		todo.Subject = sj
		todo.Description = des

		resp = append(resp, todo)

		if err != nil {
			log.Fatal(err)
		}
	}

	// Render templates
	templates.Execute(res, resp)
}

func delete(res http.ResponseWriter, req *http.Request) {
	fmt.Println("begin delete")

	// Get key from url
	key := req.URL.Query().Get("ID")

	// Delete from db
	_, err = db.Exec(fmt.Sprintf("DELETE FROM todo WHERE id=%s", key))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted")

	// Redirect to main page
	http.Redirect(res, req, "/", 301)
}

func dbConn() error {
	db, err = sql.Open("postgres", "postgres://zecolkej:KwZAmgZmsXEN_xPhnzWEsNmNRz2rNlJl@john.db.elephantsql.com:5432/zecolkej")
	if err != nil {
		return err
	}
	return nil
}
