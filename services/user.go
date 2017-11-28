package services

import (
	"github.com/SausageRocketDatingService/SausageRocketApi/models"
	"github.com/SausageRocketDatingService/SausageRocketApi/repositories"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles some of the CRUID operations of a user model.
type UserService interface {
	GetAll() ([]models.User, error)
	GetByEmail(email string) (models.User, error)
	Create(user models.User) error
	DeleteByID(id int64) error
}

// NewUserService returns a user service.
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

type userService struct {
	repo repositories.UserRepository
}

// GetAll returns all users
func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.SelectMany()
}

// GetByEmail returns a user on its email
func (s *userService) GetByEmail(email string) (models.User, error) {
	return s.repo.Select(email)
}

// Create a new user
func (s *userService) Create(user models.User) error {
	var hashedPassword string
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	hashedPassword = string(hashedBytes)
	user.Password = hashedPassword
	return s.repo.Insert(user)
}

func (s *userService) DeleteByID(id int64) error {
	return s.repo.Delete(id)
}
