package models

import "time"

type Borrow struct {
	ID uint64 `gorm:"primary_key;auto_increment;unique" json:"id"`
	UserID uint64 `gorm:"not null" json:"user_id"`
	User User `json:"user"`
	BookID uint64 `gorm:"not null" json:"book_id"`
	Book Book `json:"book"`
	BorrowedAt time.Time `json:"borrowed_at"`
}
