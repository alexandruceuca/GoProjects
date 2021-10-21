package main

import (
	"fmt"
	"html/template"
	"net/http"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Name     string
	DBStatus bool
}

func main() {

	templates := template.Must(template.ParseFiles("templates/index.html"))
	//create a database
	db, _ := sql.Open("sqlite3", "dev.db")

	//handeler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "GoProject"}

		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}

		p.DBStatus = db.Ping() == nil

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//db.Close() close the database
	})
	//web server
	//fmt.Println to print if there are any erroes
	fmt.Println(http.ListenAndServe(":8080", nil))
}
