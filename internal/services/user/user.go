package user

import (
	"kc-take-home-test/config"
	"kc-take-home-test/internal/constant"
	"kc-take-home-test/internal/models"
)

func (s *UserService) RegisterUser(nama, nik, noHP string) (result models.User, err error) {
	_, err = s.user.CheckExistingUser(nik, noHP)
	if err != nil {
		if err.Error() != constant.UserNotFound.Error() {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"nama":  nama,
				"nik":   nik,
				"no_hp": noHP,
			})
			return result, err
		}
	}

	newUser := models.User{
		Nama:  nama,
		NIK:   nik,
		NoHP:  noHP,
		Saldo: 0,
	}

	err = s.user.CreateUser(newUser)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"nama":  nama,
			"nik":   nik,
			"no_hp": noHP,
		})
		return result, err
	}

	user, err := s.user.CheckExistingUser(nik, noHP)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"nama":  nama,
			"nik":   nik,
			"no_hp": noHP,
		})
		return result, err
	}

	return user, nil
}
