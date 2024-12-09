package main

import (
	"2/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	books models.BookModel
}

func main() {

	db, err := sql.Open("mysql", "username@tcp(127.0.0.1:3306)/crud")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		books: models.BookModel{DB: db},
	}

	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)

}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {

	rows, err := env.books.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
	}

	for _, row := range rows {
		fmt.Fprintf(w, "%s,%s,%s,%f", row.Isbn, row.Title, row.Author, row.Price)
	}

}
