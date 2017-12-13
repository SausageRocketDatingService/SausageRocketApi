package services

import (
	"github.com/SausageRocketDatingService/SausageRocketApi/models"
	"github.com/SausageRocketDatingService/SausageRocketApi/repositories"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles some of the CRUID operations of a user model.
type UserService interface {
	GetAll() ([]models.User, error)
	GetByID(id int64) (models.User, error)
	Create(user models.User) error
	DeleteByID(id int64) error
	GetByUsernameAndPassword(username string, password string) (models.User, bool, error)
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

// GetByID returns a user on its id
func (s *userService) GetByID(id int64) (models.User, error) {
	return s.repo.Select(id)
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

func (s *userService) GetByUsernameAndPassword(username string, password string) (models.User, bool, error) {
	// make selectBy function in repository
	// add jwt
	// return isFound if it's found
	return models.User{}, false, nil
}
