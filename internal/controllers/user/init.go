package user

import (
	"kc-take-home-test/internal/services/user"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	RegisterUser(c echo.Context) (err error)
}

type UserController struct {
	user user.IUserService
}

func InitUserController(user user.IUserService) IUserController {
	return &UserController{
		user: user,
	}
}
