package usecases

import (
	"hexagonal02/domain/entities"
	"hexagonal02/domain/port"
	"hexagonal02/infraestructure/controllers/dto"
)

type UserService struct {
	UserRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}
func (us *UserService) CreateUser(user *entities.User) (*dto.ResponseError,error) {
	var errores []string

	if user.Name == "" {
		errores = append(errores, "Name is required")
	}
	if user.Email == "" {
        errores = append(errores, "Email is required")
    }
	if user.Password == "" {
        errores = append(errores, "Password is required")
    }
	if len(errores) > 0 {
		return dto.NewResponseError(false,errores),nil
	}	
	error :=us.UserRepository.Save(user)
	if error!= nil {
        return nil, error
    }
	return nil,nil 
}

func (us *UserService) UpdateUser(user *entities.User) (*dto.ResponseError,error) {
	var errores []string

	if user.ID == 0 {
		errores = append(errores, "ID is required")
	}
	if user.Name == "" {
		errores = append(errores, "Name is required")
	}
	if user.Email == "" {
        errores = append(errores, "Email is required")
    }
	if user.Password == "" {
        errores = append(errores, "Password is required")
    }
	if len(errores) > 0 {
		return dto.NewResponseError(false,errores),nil
	}	
	error :=us.UserRepository.Update(user)
	if error!= nil {
        return nil, error
    }
	return nil,nil 
}