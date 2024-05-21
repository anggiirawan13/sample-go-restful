package services

import (
	"github.com/sample-go-restful/models"
	"github.com/sample-go-restful/repositories"
)

type userServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) GetAllUsers() (*[]models.UserModel, error) {
	return s.userRepository.FindAll()
}

func (s *userServiceImpl) GetUserByID(id uint) (*models.UserModel, error) {
	return s.userRepository.FindByID(id)
}

func (s *userServiceImpl) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	return s.userRepository.Create(user)
}

func (s *userServiceImpl) UpdateUser(user *models.UserModel) (*models.UserModel, error) {
	return s.userRepository.Update(user)
}

func (s *userServiceImpl) DeleteUser(id uint) error {
	return s.userRepository.Delete(id)
}
