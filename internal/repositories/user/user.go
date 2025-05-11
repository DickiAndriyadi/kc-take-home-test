package user

import (
	"context"
	"kc-take-home-test/config"
	"kc-take-home-test/internal/constant"
	"kc-take-home-test/internal/models"
	"strings"

	"github.com/afex/hystrix-go/hystrix"
	"gorm.io/gorm"
)

func (r *UserRepository) CheckExistingUser(nik, noHP string) (result models.User, err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {
		query := r.db.Where("nik = ? OR no_hp = ?", nik, noHP).First(&result)
		err = query.Error

		if err != nil {

			if err.Error() == gorm.ErrRecordNotFound.Error() {
				err = constant.UserNotFound
			}

			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"nik":   nik,
				"no_hp": noHP,
			})
			return err
		}

		return nil
	}, nil); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"nik":   nik,
			"no_hp": noHP,
		})
		return result, err
	}

	return result, nil
}

func (r *UserRepository) CreateUser(user models.User) (err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {

		err = r.db.Create(&user).Error
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				err = constant.UserHasAlreadyExist
			}

			if strings.Contains(err.Error(), "duplicate key") {
				err = constant.NIKOrNoHPAlreadyExist
			}

			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"user": user,
			})
			return err
		}

		return nil
	}, nil); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"user": user,
		})
		return err
	}

	return nil
}
