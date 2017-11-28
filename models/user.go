package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

// User represents DB Model of User
type User struct {
	ID          int64
	Name        string
	Email       string
	Phone       string
	BirthDate   time.Time
	Password    string
	Location    string
	Images      []string
	Gender      string
	Description string

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

// CreateUserTable creates table for User Model in the DB
func CreateUserTable(engine *xorm.Engine) error {
	err := engine.Sync2(new(User))
	return err
}
