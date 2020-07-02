package middlewares

import (
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/auth"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/responses"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		next(writer, request)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := auth.TokenValid(request)
		if err != nil {
			responses.ResponseWithJSON(writer, http.StatusUnauthorized,
				responses.Response{Code: http.StatusUnauthorized, Message: "Unauthorized", Data: err})
			return
		}
		next(writer, request)
	}
}
