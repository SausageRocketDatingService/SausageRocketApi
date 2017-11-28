package config

import (
	"os"

	"github.com/SausageRocketDatingService/SausageRocketApi/models"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq" // driver for postgresql
)

// InitDB initializes postgres DB and returns engine of type xorm.Engine and error
func InitDB() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("postgres", os.Getenv("DATABASE_URL"))
	return engine, err
}

// CreateDBTables generates DB Tables for DB Models
func CreateDBTables(engine *xorm.Engine) error {
	userError := models.CreateUserTable(engine)
	if userError != nil {
		return userError
	}

	preferencesError := models.CreateUserPreferencesTable(engine)
	if preferencesError != nil {
		return preferencesError
	}

	ratingError := models.CreateUserRatingTable(engine)
	if ratingError != nil {
		return ratingError
	}

	return nil
}
