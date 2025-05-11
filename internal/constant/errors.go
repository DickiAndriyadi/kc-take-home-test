package constant

import (
	"kc-take-home-test/internal/models"
	"net/http"
)

var (

	// 404 Not Found
	UserNotFound = models.WrapError(http.StatusNotFound, "User tidak ditemukan!")

	// 400 Bad Request
	NIKOrNoHPAlreadyExist = models.WrapError(http.StatusBadRequest, "NIK atau No HP sudah digunakan!")
	UserHasAlreadyExist   = models.WrapError(http.StatusBadRequest, "User sudah ada!")
	InvalidRequest        = models.WrapError(http.StatusBadRequest, "Invalid request!")
	ZeroNominal           = models.WrapError(http.StatusBadRequest, "Nominal harus lebih dari 0!")
	LowBalance            = models.WrapError(http.StatusBadRequest, "Saldo tidak cukup!")
)

type RestResponse struct {
	Remark string      `json:"remark"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}
