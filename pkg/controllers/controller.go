package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Ashwin202/go-bookhub/pkg/models"
	"github.com/Ashwin202/go-bookhub/pkg/utills"
	"strconv"
)

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utills.ParseBody(r, &book)
	b := book.CreateBook()
	marshalledBook, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(marshalledBook)
}

var GetAllBooks = func (w http.ResponseWriter, r *http.Request)  {
	allBooks := models.GetAllBooks()
	marshalledBooks, _ := json.Marshal(allBooks)
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledBooks)
}

var GetBookById = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	bookID := vars["bookId"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("please send correct book id"))
		return
	}
	bookDetails, _ := models.GetBookById(ID)
	marshalledBook, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledBook)
}

var UpdateBookById = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	bookID := vars["bookId"]
	var newBookDetails models.Book
	utills.ParseBody(r, &newBookDetails)
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("please send correct book id"))
		return
	}
	bookDetails, db := models.GetBookById(ID)
	bookDetails.Name = newBookDetails.Name
	bookDetails.Author = newBookDetails.Author
	bookDetails.Publication = newBookDetails.Publication

	db.Save(&bookDetails)
	marshalledBook, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledBook)
}

var DeleteBookById = func (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	bookID := vars["bookId"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("please send correct book id"))
		return
	}
	bookDetails := models.DeleteBookById(ID)
	marshalledBook, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledBook)
}