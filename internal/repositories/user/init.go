package user

import (
	"kc-take-home-test/internal/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CheckExistingUser(nik, noHP string) (models.User, error)
	CreateUser(user models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
