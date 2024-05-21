package services

import (
	"github.com/sample-go-restful/models"
)

type UserService interface {
	GetAllUsers() (*[]models.UserModel, error)
	GetUserByID(id uint) (*models.UserModel, error)
	CreateUser(user *models.UserModel) (*models.UserModel, error)
	UpdateUser(user *models.UserModel) (*models.UserModel, error)
	DeleteUser(id uint) error
}
