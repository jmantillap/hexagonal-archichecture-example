package port

import "hexagonal02/domain/entities"

type UserRepository interface {
	Save(user *entities.User) error
	Update(user *entities.User) error
}
