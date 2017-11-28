package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

// UserPreferences represenets DB Model for user preferences
type UserPreferences struct {
	ID     int64
	UserID int64
	Goal   string
	Gender string
	Age    int

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

// CreateUserPreferencesTable creates table for UserPreferences Model in the DB
func CreateUserPreferencesTable(engine *xorm.Engine) error {
	err := engine.Sync2(new(UserPreferences))
	return err
}
