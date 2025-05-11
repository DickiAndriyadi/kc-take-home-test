package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama  string `json:"nama" gorm:"column:nama;not null;type:varchar(191)"`
	NIK   string `json:"nik" gorm:"column:nik;uniqueIndex:nik_type_unique;not null;type:varchar(191)"`
	NoHP  string `json:"no_hp" gorm:"column:no_hp;uniqueIndex:no_hp_type_unique;not null;type:varchar(191)"`
	Saldo int64  `json:"saldo" gorm:"column:saldo;default:0"`

	Transaksi []*Transaksi `json:"transaksi" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "user"
}

type Transaksi struct {
	gorm.Model
	UserID         uint      `json:"user_id" gorm:"column:user_id;not null"`
	NoHP           string    `json:"no_hp" gorm:"column:no_hp;not null"`
	Tipe           string    `json:"tipe" gorm:"column:tipe_transaksi;not null"`
	Nominal        int64     `json:"nominal" gorm:"column:nominal;not null"`
	WaktuTransaksi time.Time `json:"waktu_transaksi" gorm:"column:waktu_transaksi;not null"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}
