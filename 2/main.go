package main

import (
	"1/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

import _ "github.com/go-sql-driver/mysql"

type Env struct {
	db *sql.DB
}

func main() {

	db, err := sql.Open("mysql", "username@tcp(127.0.0.1:3306)/crud")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{db: db}

	http.HandleFunc("/books", env.handleAllBooks)
	http.ListenAndServe(":3000", nil)

}

func (env *Env) handleAllBooks(w http.ResponseWriter, r *http.Request) {

	bks, err := models.AllBooks(env.db)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s,%s,%s,%f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}
