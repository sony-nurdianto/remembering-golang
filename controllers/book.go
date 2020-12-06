package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sony-nurdianto/remembering-golang/utils"

	"github.com/sony-nurdianto/remembering-golang/models"
	"github.com/sony-nurdianto/remembering-golang/repository/bookrepository"
)

//Controller struct
type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//GetBooks Controller
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		books = []models.Book{}
		bookRepo := bookrepository.BookRepository{}

		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Messasge = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSucces(w, books)
	}
}

//GetBook Get single book
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookrepository.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Messasge = "Not Found"

				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSucces(w, book)
	}
}

//AddBook controler
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Messasge = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookrepository.BookRepository{}
		book, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Messasge = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSucces(w, book)

	}
}

//UpdateBook Controller to Handle Update Request
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Messasge = "All fields are required"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookrepository.BookRepository{}
		rowsUpdate, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Messasge = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSucces(w, rowsUpdate)
	}

}

//RemoveBook Controller to Handle Deleted Book
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)

		bookRepo := bookrepository.BookRepository{}
		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := bookRepo.RemoveBook(db, id)

		if err != nil {
			error.Messasge = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Messasge = "Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSucces(w, rowsDeleted)
	}
}
