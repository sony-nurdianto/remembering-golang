package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/remembering-golang.git/controllers"
	"github.com/subosito/gotenv"

	"github.com/sony-nurdianto/remembering-golang.git/driver"
	"github.com/sony-nurdianto/remembering-golang.git/models"
)

var book []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()
	control := controllers.Controller{}

	r := mux.NewRouter()

	r.HandleFunc("/books", control.GetBooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", control.GetBook(db)).Methods("GET")
	r.HandleFunc("/addBook", control.AddBook(db)).Methods("POST")
	r.HandleFunc("/updateBook", control.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/removeBook/{id}", control.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r)))
}
