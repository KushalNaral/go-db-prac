package main

import (
	"1/models"
	"fmt"
	"log"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

func main() {

	err := models.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {

	bks, err := models.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s,%s,%s,%f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
