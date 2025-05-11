package config

import (
	"kc-take-home-test/internal/models"
	"log"
)

func MigrateDB() {
	err := DB.Debug().AutoMigrate(&models.User{}, &models.Transaksi{})
	if err != nil {
		log.Fatalf("Database migration failed: %s", err)
	}

	// DB.Exec("ALTER TABLE transaksi ALTER COLUMN no_hp TYPE VARCHAR(50);")

	log.Println("Database migrated successfully!")
}
