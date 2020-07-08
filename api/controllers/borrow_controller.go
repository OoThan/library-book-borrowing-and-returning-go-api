package controllers

import (
	"fmt"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/auth"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/models"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/responses"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (server *Server) Take(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusBadRequest,
			responses.Response{Code: http.StatusBadRequest, Message: "Bad Request", Data: err})
		return
	}
	uid, err := auth.ExtractUserTokenID(r)
	if err != nil {
		fmt.Println(uid)
		responses.ResponseWithJSON(w, http.StatusUnauthorized,
			responses.Response{Code: http.StatusUnauthorized, Message: "Unauthorized", Data: err})
		return
	}
	var userBooksCount uint64
	borrows := []models.Borrow{}
	err = server.DB.Debug().Model(&models.Borrow{}).Where("user_id = ?", uid).Take(&borrows).Count(&userBooksCount).Error
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			 responses.Response{Code: http.StatusInternalServerError, Message: "Server Error", Data: err})
		return
	}
	if userBooksCount < 5 {
		var bookExistCount uint64
		bookExists := []models.Borrow{}
		err = server.DB.Debug().Model(&models.Borrow{}).Where("user_id = ? and book_id = ?", uid, bid).Take(&bookExists).Count(&bookExistCount).Error
		if err != nil{
			responses.ResponseWithJSON(w, http.StatusInternalServerError,
				responses.Response{Code: http.StatusInternalServerError, Message: "UserID and BookID Server Error", Data: err})
			return
		}
		if bookExistCount > 1{
			responses.ResponseWithJSON(w, http.StatusInternalServerError,
				responses.Response{Code: http.StatusInternalServerError, Message: "Book is already borrowed!", Data: bookExists})
			return
		}
		user := models.User{}
		fmt.Println(uid)
		err = server.DB.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error
		if err != nil {
			responses.ResponseWithJSON(w, http.StatusInternalServerError,
				responses.Response{Code: http.StatusInternalServerError, Message: "User ID Server Error", Data: err})
			return
		}
		book := models.Book{}
		fmt.Println(bid)
		err = server.DB.Debug().Model(&models.Book{}).Where("id = ?", bid).Take(&book).Error
		if err != nil {
			responses.ResponseWithJSON(w, http.StatusInternalServerError,
				responses.Response{Code: http.StatusInternalServerError, Message: "Book Id Server Error", Data: err})
			return
		}
		//err = server.DB.Debug().Model(&models.Borrow{}).Where("user_").Preloads()

		borrow := models.Borrow{
			UserID: uid,
			User: user,
			BookID: bid,
			Book: book,
			BorrowedAt: time.Now(),
		}
		err = server.DB.Debug().Model(&models.Book{}).Create(&borrow).Error
		if err != nil {
			responses.ResponseWithJSON(w, http.StatusInternalServerError,
				responses.Response{Code: http.StatusInternalServerError, Message: "Borrow Create Server Error", Data: err})
			return
		}
		/*err = server.DB.Debug().Model(&models.Borrow{}).Where("user_id = ? and book_id = ?", uid, bid).Preload("Borrow.User").Preload("Borrow.Book").Take(&borrow).Error
		if err != nil {
			responses.ResponseWithJSON(w, http.StatusInternalServerError,
				responses.Response{Code: http.StatusInternalServerError, Message: "Preload Error", Data: err})
			return
		}*/
		responses.ResponseWithJSON(w, http.StatusCreated,
			responses.Response{Code: http.StatusCreated, Message: "Borrow Created", Data: borrow})
	} else if userBooksCount >= 5 {
		responses.ResponseWithJSON(w, http.StatusUnprocessableEntity,
			responses.Response{Code: http.StatusUnprocessableEntity, Message: "You have already borrowed 5 books so can't borrow any books.", Data: userBooksCount})
		return
	}
}

func (server *Server) Give(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusBadRequest,
			responses.Response{Code: http.StatusBadRequest, Message: "Bad Request", Data: err})
		return
	}
	uid, err := auth.ExtractUserTokenID(r)
	if err != nil {
		fmt.Println(uid)
		responses.ResponseWithJSON(w, http.StatusUnauthorized,
			responses.Response{Code: http.StatusUnauthorized, Message: "Unauthorized", Data: err})
		return
	}
	err = server.DB.Debug().Model(&models.Borrow{}).Where("user_id = ? and book_id = ?", uid, bid).Take(&models.Borrow{}).Delete(&models.Borrow{}).Error
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Borrow Delete Server Error", Data: err})
		return
	}
	responses.ResponseWithJSON(w, http.StatusNoContent,
		responses.Response{Code: http.StatusNoContent, Message: "Borrow Delete Successfully", Data: err})
}

func (server *Server) Borrow_Count(w http.ResponseWriter, r *http.Request) {
	var count uint64
	borrows := []models.Borrow{}
	err := server.DB.Debug().Model(&models.Borrow{}).Take(&borrows).Count(&count).Error
	if err != nil {
		responses.ResponseWithJSON(w, http.StatusInternalServerError,
			responses.Response{Code: http.StatusInternalServerError, Message: "Borrow Count Server Error", Data: err})
		return
	}
	responses.ResponseWithJSON(w, http.StatusOK,
		responses.Response{Code: http.StatusOK, Message: "Total Borrower Count", Data: count})
}
