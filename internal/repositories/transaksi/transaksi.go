package transaksi

import (
	"context"
	"kc-take-home-test/config"
	"kc-take-home-test/internal/constant"
	"kc-take-home-test/internal/models"

	"github.com/afex/hystrix-go/hystrix"
	"gorm.io/gorm"
)

func (r *TransaksiRepository) GetUserByNoHP(noHP string) (result models.User, err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {
		if err := r.db.Preload("Transaksi").Where("no_hp = ?", noHP).First(&result).Error; err != nil {

			if err.Error() == gorm.ErrRecordNotFound.Error() {
				err = constant.UserNotFound
			}

			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"no_hp": noHP,
			})
			return err
		}

		return nil
	}, nil); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp": noHP,
		})
		return result, err
	}

	return result, nil
}

func (j *TransaksiRepository) UpdateSaldoUser(user models.User, nominal int64) (err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {

		user.Saldo += nominal
		if err := j.db.Where("id = ?", user.ID).Updates(&user).Error; err != nil {
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

func (r *TransaksiRepository) CreateTransaksi(transaksi models.Transaksi) (err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {

		err = r.db.Create(&transaksi).Error
		if err != nil {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"transaksi": transaksi,
			})
			return err
		}

		return nil
	}, nil); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"transaksi": transaksi,
		})
		return err
	}

	return nil
}
