# library-book-borrowing-and-returning-go-api

> A complete example of a REST-ful Library System API server in written Go (Golang)

## Quick Start

``` bash
# install mux router
go get -u github.com/gorilla/mux

# install gorm for MySQL Database
go get -u github.com/jinzhu/gorm

# install jwt-go for user token authentication
go get -u github.com/dgrijalva/jwt-go

# install checkmail to validate email
go get -u github.com/badoux/checkmail

# install crypto to encrypt and decrypt user password
go get -u golang.org/x/crypto

# install godotenv to get .env file from os
go get -u github.com/joho/godotenv
```

```bash
go run main.go
    (or)
go build
./library-book-borrowing-and-returning-go-api
```

##Project Directory
```.
├── main.go                                 // entry file
├── .env                                    // for mysql and jwt-go
├── books.json                              // dataset for books
├── go.mod                                  // go module
├── go.sum                                  // go module status
└── api
      ├── auth
      |       └── token.go                    // generate user token and validate user id
      ├── controllers
      |       ├── base.go                     // initialize server
      |       ├── book_controller.go          // control book model
      |       ├── borrow_controller.go        // control borrow model
      |       ├── home_controller.go          // for home page
      |       ├── routes.go                   // routes initialize
      |       └── user_controller.go          // control user model
      ├── data
      |       └── books.json      
      ├── middlewares                        
      |       └── middleware.go               // for request and respond 
      ├── models
      |       ├── book.go                     // boook model
      |       ├── borrow.go                   // borrow model
      |       └── user.go                     // user model
      ├── responses
      |       └── utils.go                    // for JSON responses
      ├── seed     
      |       └── sesder.go                   // to seed dataset into mysql
      ├── utils                       
      |       └── format         
      |               └── format_error.go     // format error     
      └── server.go                           // server start 

```

## End Points

#### Home
```
GET /
```
#### User
```
POST /register
# Request simple
{
	"user_name": "Than Hla",
	"email": "thanhla@gmail.com",
	"password": "thanhla"
}

POST /login
# Request simple
{
	"email": "thanhla@gmail.com",
	"password": "thanhla"
}

GET /users
GET /user_count
```
#### Book
```
GET /books

POST /books
# Create Book Request Simple
{
    "_id": 1,
    "title": "Unlocking Android",
    "isbn": "1933988673",
    "pageCount": 416,
    "publishedDate": {
      "$date": "2009-04-01T00:00:00.000-0700"
    },
    "thumbnailUrl": "https://s3.amazonaws.com/AKIAJC5RLADLUMVRPFDQ.book-thumb-images/ableson.jpg",
    "shortDescription": "Unlocking Android: A Developer's Guide provides concise, hands-on instruction for the Android operating system and development tools. This book teaches important architectural concepts in a straightforward writing style and builds on this with practical and useful examples throughout.",
    "longDescription": "Android is an open source mobile phone platform based on the Linux operating system and developed by the Open Handset Alliance, a consortium of over 30 hardware, software and telecom companies that focus on open standards for mobile devices. Led by search giant, Google, Android is designed to deliver a better and more open and cost effective mobile experience.    Unlocking Android: A Developer's Guide provides concise, hands-on instruction for the Android operating system and development tools. This book teaches important architectural concepts in a straightforward writing style and builds on this with practical and useful examples throughout. Based on his mobile development experience and his deep knowledge of the arcane Android technical documentation, the author conveys the know-how you need to develop practical applications that build upon or replace any of Androids features, however small.    Unlocking Android: A Developer's Guide prepares the reader to embrace the platform in easy-to-understand language and builds on this foundation with re-usable Java code examples. It is ideal for corporate and hobbyists alike who have an interest, or a mandate, to deliver software functionality for cell phones.    WHAT'S INSIDE:        * Android's place in the market      * Using the Eclipse environment for Android development      * The Intents - how and why they are used      * Application classes:            o Activity            o Service            o IntentReceiver       * User interface design      * Using the ContentProvider to manage data      * Persisting data with the SQLite database      * Networking examples      * Telephony applications      * Notification methods      * OpenGL, animation & multimedia      * Sample Applications  ",
    "status": "PUBLISH",
    "authors": [
      "W. Frank Ableson",
      "Charlie Collins",
      "Robi Sen"
    ],
    "categories": [
      "Open Source",
      "Mobile"
    ]
  }

GET /book/{id}

PUT /book/{id}
# Update Book Request Simple
{
    "_id": 1,
    "title": "Unlocking Android",
    "isbn": "1933988673",
    "pageCount": 416,
    "publishedDate": {
      "$date": "2009-04-01T00:00:00.000-0700"
    },
    "thumbnailUrl": "https://s3.amazonaws.com/AKIAJC5RLADLUMVRPFDQ.book-thumb-images/ableson.jpg",
    "shortDescription": "Unlocking Android: A Developer's Guide provides concise, hands-on instruction for the Android operating system and development tools. This book teaches important architectural concepts in a straightforward writing style and builds on this with practical and useful examples throughout.",
    "longDescription": "Android is an open source mobile phone platform based on the Linux operating system and developed by the Open Handset Alliance, a consortium of over 30 hardware, software and telecom companies that focus on open standards for mobile devices. Led by search giant, Google, Android is designed to deliver a better and more open and cost effective mobile experience.    Unlocking Android: A Developer's Guide provides concise, hands-on instruction for the Android operating system and development tools. This book teaches important architectural concepts in a straightforward writing style and builds on this with practical and useful examples throughout. Based on his mobile development experience and his deep knowledge of the arcane Android technical documentation, the author conveys the know-how you need to develop practical applications that build upon or replace any of Androids features, however small.    Unlocking Android: A Developer's Guide prepares the reader to embrace the platform in easy-to-understand language and builds on this foundation with re-usable Java code examples. It is ideal for corporate and hobbyists alike who have an interest, or a mandate, to deliver software functionality for cell phones.    WHAT'S INSIDE:        * Android's place in the market      * Using the Eclipse environment for Android development      * The Intents - how and why they are used      * Application classes:            o Activity            o Service            o IntentReceiver       * User interface design      * Using the ContentProvider to manage data      * Persisting data with the SQLite database      * Networking examples      * Telephony applications      * Notification methods      * OpenGL, animation & multimedia      * Sample Applications  ",
    "status": "PUBLISH",
    "authors": [
      "W. Frank Ableson",
      "Charlie Collins",
      "Robi Sen"
    ],
    "categories": [
      "Open Source",
      "Mobile"
    ]
  }

DELETE /book/{id}

GET /book_count
```
#### Borrow
```
GET /borrow_count

POST /book_take/{id}

DELETE /book_give/{id}
```