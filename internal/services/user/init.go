package user

import (
	"kc-take-home-test/internal/models"
	"kc-take-home-test/internal/repositories/user"
)

type IUserService interface {
	RegisterUser(nama, nik, noHP string) (result models.User, err error)
}

type UserService struct {
	user user.IUserRepository
}

func InitUserService(user user.IUserRepository) IUserService {
	return &UserService{
		user: user,
	}
}
