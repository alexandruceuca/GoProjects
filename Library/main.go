package main

import (
	"fmt"
	"html/template"
	"net/http"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"encoding/json"
)

type Page struct {
	Name     string
	DBStatus bool
}

type SearchResult struct {
	Title  string
	Author string
	Year   string
	ID     string
}

func main() {

	templates := template.Must(template.ParseFiles("templates/index.html"))
	//create a database
	db, _ := sql.Open("sqlite3", "dev.db")

	//handeler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "GoProject"}
		//query
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		//status of database
		p.DBStatus = db.Ping() == nil

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//db.Close() close the database
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		//dummy data
		results := []SearchResult{
			SearchResult{"Moby Dick", "Herman Melville", "1851", "222222"},
			SearchResult{"The Adventures of Huckleberry Finn", "Mark Twain", "1884", "444444"},
			SearchResult{"The Catcher in the Rye", "JD Salinger", "1951", "333333"},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	})

	//web server
	//fmt.Println to print if there are any erroes
	fmt.Println(http.ListenAndServe(":8080", nil))
}
