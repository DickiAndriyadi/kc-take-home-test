package user

import (
	"kc-take-home-test/config"
	"kc-take-home-test/internal/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ctrl *UserController) RegisterUser(c echo.Context) (err error) {
	var req struct {
		Nama string `json:"nama"`
		NIK  string `json:"nik"`
		NoHP string `json:"no_hp"`
	}
	if err := c.Bind(&req); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"req": req,
		})
		return c.JSON(http.StatusBadRequest, constant.RestResponse{
			Remark: err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
	}

	result, err := ctrl.user.RegisterUser(req.Nama, req.NIK, req.NoHP)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"req": req,
		})
		return c.JSON(http.StatusBadRequest, constant.RestResponse{
			Remark: err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
	}

	return c.JSON(http.StatusOK, constant.RestResponse{
		Remark: "User berhasil dibuat!",
		Code:   http.StatusOK,
		Data:   result,
	})
}
