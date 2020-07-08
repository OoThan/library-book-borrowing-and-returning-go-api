package seed

import (
	"encoding/json"
	"fmt"
	"github.com/OoThan/library-book-borrowing-and-returning-go-api/api/models"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"
)

var users = []models.User{
	models.User{
		UserName: "DawMya",
		Email:    "mya@gmail.com",
		Password: "user1",
	},
	models.User{
		UserName: "U Ba",
		Email:    "uba@gmail.com",
		Password: "user2",
	},
	models.User{
		UserName: "Daw Sein",
		Email:    "dawsein@gmail.com",
		Password: "user3",
	},
	models.User{
		UserName: "Maung Maung",
		Email:    "maungmanug@mail.com",
		Password: "user4",
	},
	models.User{
		UserName: "Chaw Chaw",
		Email:    "chawchaw@gmail.com",
		Password: "user5",
	},
}

func Load(db *gorm.DB) {
	//Parsing data from books.json
	jsonFile, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var books []models.Book
	//var books []models.Book
	byteBookValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println("Successfully open JSON file!")
	err = json.Unmarshal(byteBookValue, &books)
	if err != nil {
		fmt.Println(err)
		//		fmt.Println(len(booksJSON), " ", len(books))
	}
	/*fmt.Println(len(booksJSON), " ", len(books))
	for i := 0; i < len(booksJSON)-1; i++ {
		fmt.Println(len(booksJSON), " ", books[i].ID)
		books[i].ID = booksJSON[i].ID
		books[i].Title = booksJSON[i].Title
		books[i].ISBN = booksJSON[i].ISBN
		books[i].PageCount = booksJSON[i].PageCount
		books[i].ThumbnailURL = booksJSON[i].ThumbnailURL
		books[i].ShortDescription = booksJSON[i].ShortDescription
		books[i].LongDescription = booksJSON[i].LongDescription
		books[i].Status = booksJSON[i].Status
	}*/

	err = db.Debug().DropTableIfExists(&models.Borrow{}, &models.User{}, &models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Book{}, &models.Borrow{}).Error
	if err != nil {
		log.Fatalf("cannot migrate tables: %v", err)
	}
	err = db.Debug().Model(&models.Borrow{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("atttaching user_id foreign key error: %v ", err)
	}
	err = db.Debug().Model(&models.Borrow{}).AddForeignKey("book_id", "books(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("atttaching book_id foreign key error: %v ", err)
	}
	for i, _ := range books {
		err = db.Debug().Model(&models.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed books table: %v", err)
		}
	}
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
