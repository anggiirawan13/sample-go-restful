package repositories

import (
	"github.com/sample-go-restful/models"
)

type UserRepository interface {
	FindAll() (*[]models.UserModel, error)
	FindByID(id uint) (*models.UserModel, error)
	Create(user *models.UserModel) (*models.UserModel, error)
	Update(user *models.UserModel) (*models.UserModel, error)
	Delete(id uint) error
}
