package controllers

import "github.com/OoThan/library-book-borrowing-and-returning-go-api/api/middlewares"

func (server *Server) initializeRoutes() {
	//System
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	//User
	server.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(server.RegisterUser)).Methods("POST")
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetAllUsers)).Methods("GET")
	server.Router.HandleFunc("/user_count", middlewares.SetMiddlewareJSON(server.GetUsersCount)).Methods("GET")

	//Book
	server.Router.HandleFunc("/books", middlewares.SetMiddlewareJSON(server.CreateBook)).Methods("POST")
	server.Router.HandleFunc("/books", middlewares.SetMiddlewareJSON(server.GetAllBooks)).Methods("GET")
	server.Router.HandleFunc("/book/{id}", middlewares.SetMiddlewareJSON(server.GetBookByID)).Methods("GET")
	server.Router.HandleFunc("/book/{id}", middlewares.SetMiddlewareJSON(server.UpdateBook)).Methods("PUT")
	server.Router.HandleFunc("/book/{id}", middlewares.SetMiddlewareJSON(server.DeleteBook)).Methods("DELETE")
	server.Router.HandleFunc("/book_count", middlewares.SetMiddlewareJSON(server.GetBooksCount)).Methods("GET")

	//Borrow
	server.Router.HandleFunc("/borrow_count", middlewares.SetMiddlewareJSON(server.Borrow_Count)).Methods("GET")
	server.Router.HandleFunc("/book_take/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Take))).Methods("POST")
	server.Router.HandleFunc("/book_give/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Give))).Methods("DELETE")
}
