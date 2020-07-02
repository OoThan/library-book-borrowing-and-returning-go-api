package models

import (
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserName string `gorm:"size:255;unique;not null" json:"user_name"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type JwtToken struct {
	Token string `json:"token"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Prepare() {
	user.ID = 0
	user.UserName = html.EscapeString(strings.TrimSpace(user.UserName))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "register":
		if user.UserName == "" {
			return errors.New("Required NickName! ")
		}
		if user.Email == "" {
			return errors.New("Required Email! ")
		}
		if user.Password == "" {
			return errors.New("Required Password! ")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email! ")
		}
		return nil

	case "login":
		if user.Email == "" {
			return errors.New("Required Email! ")
		}
		if user.Password == "" {
			return errors.New("Required Password! ")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email! ")
		}
		return nil

	default:
		if user.UserName == "" {
			return errors.New("Required NickName! ")
		}
		if user.Email == "" {
			return errors.New("Required Email! ")
		}
		if user.Password == "" {
			return errors.New("Required Password! ")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email! ")
		}
		return nil
	}
}

func (user *User) Register(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) Login(db *gorm.DB, email, password string) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	err = VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found! ")
	}
	fmt.Println(user)
	return user, nil
}

func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func (user *User) FindUsersCount(db *gorm.DB) (int64, error) {
	var err error
	var count int64
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
