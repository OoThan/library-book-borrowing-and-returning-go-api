package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

/*type Book struct {
	ID uint64 `json:"_id"`
	Title string `json:"title"`
	ISBN string `json:"isbn"`
	PageCount uint32 `json:"pageCount"`
	ThumbnailURL string `json:"thumbnailUrl"`
	ShortDescription string `gorm:"size:16777215" json:"shortDescription" binding:"required"`
	LongDescription string `gorm:"size:16777215" json:"longDescription" binding:"required"`
	Status string `json:"status"`
}*/

type Book struct {
	ID uint64 `gorm:"unique" json:"_id"`
	Title string `json:"title"`
	ISBN string `json:"isbn"`
	PageCount uint32 `json:"pageCount"`
	PublishedDate PublisherDate `json:"publishedDate"`
	ThumbnailURL string `json:"thumbnailUrl"`
	ShortDescription string `gorm:"size:16777215" json:"shortDescription" binding:"required"`
	LongDescription string `gorm:"size:16777215" json:"longDescription" binding:"required"`
	Status string `json:"status"`
	Author []Author `json:"authors"`
	Category []Category `json:"categories"`
}

type Author struct {

}

type Category struct {

}

type PublisherDate struct {

}

func (book *Book) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if book.Title == "" {
			return errors.New("Required Title ")
		}
		if book.ISBN == "" {
			return errors.New("Required ISBN ")
		}
		if book.ThumbnailURL == "" {
			return errors.New("Required ThumbnailUrl ")
		}
		if book.ShortDescription == "" {
			return errors.New("Required ShortDescription ")
		}
		if book.LongDescription == "" {
			return errors.New("Required LongDescription ")
		}
		if book.Status == "" {
			return errors.New("Required Status ")
		}
		return nil

	default:
		return nil
	}
}

func (book *Book) SaveBook(db *gorm.DB) (*Book, error) {
	var err error
	err = db.Debug().Create(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, err
}

func (book *Book) FindAllBooks(db *gorm.DB) (*[]Book, error) {
	var err error
	books := []Book{}
	err = db.Debug().Model(&Book{}).Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &books, err
}

func (book *Book) FindBooksCount(db *gorm.DB) (int64, error) {
	var err error
	var count int64
	books := []Book{}
	err = db.Debug().Model(&Book{}).Find(&books).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (book *Book) FindBookByID(db *gorm.DB, bid uint64) (*Book, error) {
	var err error
	err = db.Debug().Model(&Book{}).Where("id = ?", bid).Take(&book).Error
	if err != nil {
		return &Book{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Book{}, errors.New("Book Not Found! ")
	}
	return book, err
}

func (book *Book) UpdateBook(db *gorm.DB, bid uint64) (*Book, error) {
	db = db.Debug().Model(&Book{}).Where("id = ?", bid).Take(&Book{}).UpdateColumns(
		map[string]interface{}{
			"title": book.Title,
			"isbn": book.ISBN,
			"pageCount": book.PageCount,
			"thumbnailUrl": book.ThumbnailURL,
			"shortDescription": book.ShortDescription,
			"longDescription": book.LongDescription,
			"status": book.Status,
		})
	if db.Error != nil {
		return &Book{}, db.Error
	}
	err := db.Debug().Model(&Book{}).Where("id = ?", bid).Take(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}

func (book *Book) DeleteBook(db *gorm.DB, bid uint64) (int64, error) {
	db = db.Debug().Model(&Book{}).Where("id = ?", bid).Take(&Book{}).Delete(&Book{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
