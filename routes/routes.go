package routes

import (
	"kc-take-home-test/internal/controllers/transaksi"
	"kc-take-home-test/internal/controllers/user"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userCtrl user.IUserController, transaksiCtrl transaksi.ITransaksiController) {
	api := e.Group("/api")
	api.POST("/daftar", userCtrl.RegisterUser)
	api.POST("/tabung", transaksiCtrl.Tabung)
	api.POST("/tarik", transaksiCtrl.Tarik)
	api.GET("/saldo/:no_hp", transaksiCtrl.GetSaldo)
}
