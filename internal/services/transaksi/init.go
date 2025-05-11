package transaksi

import (
	"kc-take-home-test/internal/models"
	"kc-take-home-test/internal/repositories/transaksi"
)

type ITransaksiService interface {
	Tabung(noHP string, nominal int64) (models.User, error)
	Tarik(noHP string, nominal int64) (models.User, error)
	GetSaldo(noHP string) (user models.User, err error)
}

type TransaksiService struct {
	transaksi transaksi.ITransaksiRepository
}

func InitTransaksiService(transaksi transaksi.ITransaksiRepository) ITransaksiService {
	return &TransaksiService{
		transaksi: transaksi,
	}
}
