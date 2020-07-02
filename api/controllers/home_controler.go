package controllers

import (
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.ResponseWithJSON(w, http.StatusOK,
		responses.Response{Code: http.StatusOK, Message: "Welcome to MY LIBRARY SYSTEM!"})
}
