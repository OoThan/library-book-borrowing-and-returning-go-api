package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/auth"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/models"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/responses"
	"io/ioutil"
	"net/http"
)

func (server *Server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	user.Prepare()
	err = user.Validate("register")
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	userCreated, err := user.Register(server.DB)
	if err != nil {
		//formattedError := format.FormatError(err.Error())
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error! ", Data: err})
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.ResponseWithJSON(w, http.StatusCreated,
		responses.Response{Code: http.StatusCreated, Message: "Registered Success", Data: userCreated})
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	err = user.Validate("login")
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "Unprocessable", Data: err})
		return
	}
	fmt.Println(user)
	userLogin, err := user.Login(server.DB, user.Email, user.Password)
	fmt.Println(userLogin, err)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error! ", Data: err})
		return
	}
	token, _ := auth.GenerateToken(userLogin)
	responses.ResponseWithJSON(w, http.StatusOK,
		responses.Response{Code: http.StatusOK, Message: "Signed!", Data: models.JwtToken{Token: token}})
}

func (server *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error!", Data: err})
		return
	}
	responses.ResponseWithJSON(w, http.StatusOK,
		responses.Response{Code: http.StatusOK, Message: "All Users", Data: users})
}

func (server *Server) GetUsersCount(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	count, err := user.FindUsersCount(server.DB)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Server Error!", Data: err})
		return
	}
	responses.ResponseWithJSON(w, http.StatusOK,
		responses.Response{Code: http.StatusOK, Message: "All Users Count", Data: count})
}
