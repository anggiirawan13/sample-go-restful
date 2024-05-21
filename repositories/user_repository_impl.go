package repositories

import (
	"github.com/sample-go-restful/models"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) FindAll() ([]models.UserModel, error) {
	var users []models.UserModel
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *userRepositoryImpl) FindByID(id uint) (models.UserModel, error) {
	var user models.UserModel
	err := r.DB.First(&user, id).Error
	return user, err
}

func (r *userRepositoryImpl) Create(user models.UserModel) (models.UserModel, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *userRepositoryImpl) Update(user models.UserModel) (models.UserModel, error) {
	err := r.DB.Save(&user).Error
	return user, err
}

func (r *userRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&models.UserModel{}, id).Error
}
