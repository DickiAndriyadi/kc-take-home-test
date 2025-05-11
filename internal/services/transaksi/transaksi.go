package transaksi

import (
	"kc-take-home-test/config"
	"kc-take-home-test/internal/constant"
	"kc-take-home-test/internal/models"
	"time"
)

func (s *TransaksiService) Tabung(noHP string, nominal int64) (result models.User, err error) {
	if nominal <= int64(0) {
		err = constant.ZeroNominal
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	user, err := s.transaksi.GetUserByNoHP(noHP)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	err = s.transaksi.UpdateSaldoUser(user, nominal)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	transaksi := models.Transaksi{
		UserID:         user.ID,
		NoHP:           noHP,
		Tipe:           "tabung",
		Nominal:        nominal,
		WaktuTransaksi: time.Now(),
	}
	err = s.transaksi.CreateTransaksi(transaksi)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	result, err = s.transaksi.GetUserByNoHP(noHP)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	return result, nil
}

func (s *TransaksiService) Tarik(noHP string, nominal int64) (result models.User, err error) {
	if nominal <= 0 {
		err = constant.ZeroNominal
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	user, err := s.transaksi.GetUserByNoHP(noHP)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	if user.Saldo < nominal {
		err = constant.LowBalance
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	err = s.transaksi.UpdateSaldoUser(user, -nominal)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	transaksi := models.Transaksi{
		UserID:         user.ID,
		NoHP:           noHP,
		Tipe:           "tarik",
		Nominal:        -nominal,
		WaktuTransaksi: time.Now(),
	}
	err = s.transaksi.CreateTransaksi(transaksi)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	result, err = s.transaksi.GetUserByNoHP(noHP)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp":   noHP,
			"nominal": nominal,
		})
		return result, err
	}

	return result, nil
}

func (s *TransaksiService) GetSaldo(noHP string) (user models.User, err error) {
	user, err = s.transaksi.GetUserByNoHP(noHP)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_hp": noHP,
		})
		return user, err
	}
	return user, nil
}
