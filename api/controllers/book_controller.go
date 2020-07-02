package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/models"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/responses"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	book := models.Book{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	bookCreated, err := book.SaveBook(server.DB)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error", Data: err})
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, bookCreated.ID))
	responses.ResponseWithJSON(w, http.StatusCreated,
		responses.Response{Code: http.StatusCreated, Message: "Book Created", Data: bookCreated})
}

func (server *Server) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	books, err := book.FindAllBooks(server.DB)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error", Data: err})
		return
	}
	responses.ResponseWithJSON(w, http.StatusOK,
		responses.Response{Code: http.StatusOK, Message: "All Books", Data: books})
}

func (server *Server) GetBooksCount(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	count, err := book.FindBooksCount(server.DB)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error", Data: err})
		return
	}
	responses.ResponseWithJSON(w, http.StatusOK,
		responses.Response{Code: http.StatusOK, Message: "Book Count", Data: count})
}

func (server *Server) GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusBadRequest,
			responses.Response{Code: http.StatusBadRequest, Message: "Bad Request", Data: err})
		return
	}
	book := models.Book{}
	bookGotten, err := book.FindBookByID(server.DB, uint64(bid))
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error", Data: err})
		return
	}
	responses.ResponseWithJSON(w, http.StatusOK,
		responses.Response{Code: http.StatusOK, Data: bookGotten})
}

func (server *Server) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusBadRequest,
			responses.Response{Code: http.StatusBadRequest, Message: "Bad Request", Data: err})
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	book := models.Book{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	err = book.Validate("update")
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	updatedBook, err := book.UpdateBook(server.DB, uint64(bid))
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error", Data: err})
		return
	}
	responses.ResponseWithJSON(w, http.StatusOK,
		 responses.Response{Code: http.StatusOK, Message: "Update OK", Data: updatedBook})
}

func (server *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book := models.Book{}
	bid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusBadRequest,
			responses.Response{Code: http.StatusBadRequest, Message: "Bad Request", Data: err})
		return
	}
	rowAffected, err := book.DeleteBook(server.DB, uint64(bid))
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error", Data: err})
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", bid))
	responses.ResponseWithJSON(w, http.StatusNoContent,
		responses.Response{Code: http.StatusNoContent, Message: "Deleted Successful", Data: rowAffected})
}
