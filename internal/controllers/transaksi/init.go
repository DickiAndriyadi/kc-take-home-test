package transaksi

import (
	"kc-take-home-test/internal/services/transaksi"

	"github.com/labstack/echo/v4"
)

type ITransaksiController interface {
	Tabung(c echo.Context) (err error)
	Tarik(c echo.Context) (err error)
	GetSaldo(c echo.Context) (err error)
}

type TransaksiController struct {
	transaksi transaksi.ITransaksiService
}

func InitTransaksiController(transaksi transaksi.ITransaksiService) ITransaksiController {
	return &TransaksiController{
		transaksi: transaksi,
	}
}
