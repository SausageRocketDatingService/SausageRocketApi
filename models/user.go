package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

// User represents DB Model of User
type User struct {
	ID          int64
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	BirthDate   time.Time `json:"birthDate"`
	Password    string    `json:"password"`
	Location    string    `json:"location"`
	Images      []string  `json:"images"`
	Gender      string    `json:"gender"`
	Description string    `json:"description"`

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

// CreateUserTable creates table for User Model in the DB
func CreateUserTable(engine *xorm.Engine) error {
	err := engine.Sync2(new(User))
	return err
}
