package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Book struct (Model)
type Book struct {
	Isbn      string `json:"isbn"`
	Bookname  string `json:"bookname"`
	Bookprice int    `json:"bookprice"`
}

//Init Book Var as slice Book Struct
var books []Book

var str = ""

//Init Connection String
var comn = "server=localhost\\SQLEXPRESS;user id=sa;password=k7807907;port=1433;database=bookshop;"

//Get Books
func getBooks(w http.ResponseWriter, r *http.Request) {

	db, err := sqlx.Open("mssql", comn)
	e(err)
	defer db.Close()

	books := Book{}
	tsql := fmt.Sprintf(`SELECT * FROM books`)
	rows, err := db.Queryx(tsql)

	for rows.Next() {
		err = rows.StructScan(&books)
		if err != nil {
			log.Fatalln(err)
		}
		js, err := json.Marshal(books)
		e(err)
		fmt.Println(string(js))
		//str += string(js)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}

}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sqlx.Open("mssql", comn)
	e(err)
	defer db.Close()
	var book Book
	book.Isbn = params["isbn"]
	_ = json.NewDecoder(r.Body).Decode(&book)
	tsql := fmt.Sprintf(`SELECT * FROM books WHERE isbn = '%s'`, book.Isbn)
	stmit, err := db.Queryx(tsql)

	if err != nil {
		e(err)
	}

	rows, err := db.Queryx(tsql)

	for rows.Next() {
		err = rows.StructScan(&book)
		if err != nil {
			log.Fatalln(err)
		}
		js, err := json.Marshal(book)
		e(err)
		fmt.Println(string(js))
		//str += string(js)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	}

	fmt.Print("Select Single Book")
	w.Header().Set("Server", "Select Single Book")
	w.WriteHeader(200)
	defer stmit.Close()
}

//Create Book
func createBooks(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open("mssql", comn)
	e(err)
	defer db.Close()
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	tsql := fmt.Sprintf(`INSERT INTO [dbo].[books]([bookname],[bookprice]) VALUES ('%s','%d')`, book.Bookname, book.Bookprice)
	stmit, err := db.Queryx(tsql)

	if err != nil {
		e(err)
	}
	fmt.Print("Book was Create")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Book was Create")
	w.WriteHeader(200)
	defer stmit.Close()

}

//Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sqlx.Open("mssql", comn)
	e(err)
	defer db.Close()
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.Isbn = params["isbn"]

	tsql := fmt.Sprintf(`UPDATE [dbo].[books] SET [bookname] = '%s',[bookprice] = '%d'
	WHERE [isbn] = '%s'`, book.Bookname, book.Bookprice, book.Isbn)
	updatestm, err := db.Queryx(tsql)

	if err != nil {
		e(err)
	}
	fmt.Print("Book was Update")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Book was Updated")
	w.WriteHeader(200)
	defer updatestm.Close()
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db, err := sqlx.Open("mssql", comn)
	e(err)
	defer db.Close()
	var book Book
	book.Isbn = params["isbn"]
	_ = json.NewDecoder(r.Body).Decode(&book)
	tsql := fmt.Sprintf(`DELETE FROM books WHERE isbn = '%s'`, book.Isbn)
	stmit, err := db.Queryx(tsql)

	if err != nil {
		e(err)
	}
	fmt.Print("Book was deleted")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Book was deleted")
	w.WriteHeader(200)
	defer stmit.Close()
}

func main() {
	//Innit Router
	r := mux.NewRouter()
	//Rounter Handler //End Point
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{isbn}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/book/{isbn}", updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{isbn}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
