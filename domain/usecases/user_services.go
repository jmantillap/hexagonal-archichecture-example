package usecases

import (
	"hexagonal02/domain/entities"
	"hexagonal02/domain/port"
)

type UserService struct {
	UserRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}
func (us *UserService) CreateUser(user *entities.User) error {
	return us.UserRepository.Save(user)
}
