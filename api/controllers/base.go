package controllers

import (
	"fmt"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	var err error
	DBURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
							DBUser,
							DBPassword,
							DBHost,
							DBPort,
							DBName)
	server.DB, err = gorm.Open(DBDriver, DBURI)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DBDriver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the %s database", DBDriver)
	}
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Book{}, &models.Borrow{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 9090")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
