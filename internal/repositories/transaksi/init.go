package transaksi

import (
	"kc-take-home-test/internal/models"

	"gorm.io/gorm"
)

type ITransaksiRepository interface {
	GetUserByNoHP(noHP string) (models.User, error)
	UpdateSaldoUser(user models.User, nominal int64) (err error)
	CreateTransaksi(transaksi models.Transaksi) error
}

type TransaksiRepository struct {
	db *gorm.DB
}

func InitTransaksiRepository(db *gorm.DB) ITransaksiRepository {
	return &TransaksiRepository{
		db: db,
	}
}
