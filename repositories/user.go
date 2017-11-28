package repositories

import (
	"sync"

	"github.com/go-xorm/xorm"

	"github.com/SausageRocketDatingService/SausageRocketApi/models"
)

// UserRepository handles the basic operations of a user model.
type UserRepository interface {
	SelectMany() (users []models.User, err error)
	Select(email string) (user models.User, err error)
	Insert(user models.User) (err error)
	Delete(id int64) (err error)
}

// NewUserRepository returns a new user repository,
func NewUserRepository(engine xorm.Engine) UserRepository {
	return &userRepository{engine: engine}
}

type userRepository struct {
	engine xorm.Engine
	mu     sync.RWMutex
}

// Select returns a user by email
func (r *userRepository) Select(email string) (models.User, error) {
	var user = models.User{Email: email}
	_, err := r.engine.Get(&user)
	return user, err
}

// SelectMany returns all users
func (r *userRepository) SelectMany() ([]models.User, error) {
	var users []models.User
	err := r.engine.Find(&users)
	return users, err
}

// Insert creates a new user and returns it
func (r *userRepository) Insert(user models.User) error {
	_, err := r.engine.Insert(user)
	return err
}

// Insert creates a new user and returns it
func (r *userRepository) Delete(id int64) error {
	var user models.User
	_, err := r.engine.Id(id).Delete(&user)
	return err
}
