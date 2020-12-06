package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/remembering-golang/controllers"
	"github.com/subosito/gotenv"

	"github.com/sony-nurdianto/remembering-golang/driver"
	"github.com/sony-nurdianto/remembering-golang/models"
)

var book []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hallo")
}

func main() {

	db = driver.ConnectDB()
	control := controllers.Controller{}
	port := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/books", control.GetBooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", control.GetBook(db)).Methods("GET")
	r.HandleFunc("/addBook", control.AddBook(db)).Methods("POST")
	r.HandleFunc("/updateBook", control.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/removeBook/{id}", control.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r)))
}
