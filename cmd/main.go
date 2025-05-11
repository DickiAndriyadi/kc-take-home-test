package main

import (
	"kc-take-home-test/config"
	"kc-take-home-test/routes"
	"log"

	userController "kc-take-home-test/internal/controllers/user"
	userRepo "kc-take-home-test/internal/repositories/user"
	userService "kc-take-home-test/internal/services/user"

	transaksiController "kc-take-home-test/internal/controllers/transaksi"
	transaksiRepo "kc-take-home-test/internal/repositories/transaksi"
	transaksiService "kc-take-home-test/internal/services/transaksi"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment & database
	config.ConnectDB()
	config.MigrateDB()

	// Argument parser untuk host & port
	configs := config.ParseArgs()

	// Inisialisasi repository, service, dan controller
	uRepository := userRepo.InitUserRepository(config.DB)
	uService := userService.InitUserService(uRepository)
	uController := userController.InitUserController(uService)

	tRepository := transaksiRepo.InitTransaksiRepository(config.DB)
	tService := transaksiService.InitTransaksiService(tRepository)
	tController := transaksiController.InitTransaksiController(tService)

	// Inisialisasi Echo
	e := echo.New()

	// Routing
	routes.InitRoutes(e, uController, tController)

	// Jalankan server dengan argument parser
	serverAddr := configs.Host + ":" + configs.Port
	log.Println("Server berjalan di", serverAddr)
	e.Logger.Fatal(e.Start(serverAddr))
}
