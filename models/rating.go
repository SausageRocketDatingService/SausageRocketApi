package models

import "github.com/go-xorm/xorm"

// UserRating represents DB Model of UserRating
type UserRating struct {
	ID     int64
	UserID int64
	Score  int
}

// CreateUserRatingTable creates table for UserRating Model in the DB
func CreateUserRatingTable(engine *xorm.Engine) error {
	err := engine.Sync2(new(UserRating))
	return err
}
