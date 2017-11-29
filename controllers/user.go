package controllers

import (
	"github.com/SausageRocketDatingService/SausageRocketApi/models"
	"github.com/SausageRocketDatingService/SausageRocketApi/services"
	"github.com/kataras/iris/mvc"
)

// UserController controlls routes for /user
type UserController struct {
	mvc.C
	Service services.UserService
}

// Get returns an array of users or an error
func (c *UserController) Get() ([]models.User, error) {
	users, err := c.Service.GetAll()
	return users, err
}

// GetBy returns an a user by id or an error
func (c *UserController) GetBy(id int64) (models.User, error) {
	user, err := c.Service.GetByID(id)
	return user, err
}

// Put creates new user
func (c *UserController) Put() error {
	var user models.User
	err := c.Ctx.ReadJSON(&user)
	if err != nil {
		return err
	}
	err = c.Service.Create(user)
	return err
}

// DeleteBy deletes user by id
func (c *UserController) DeleteBy(id int64) error {
	err := c.Service.DeleteByID(id)
	return err
}
