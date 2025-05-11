package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	conn := func() (err error) {
		err = godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found, using system environment variables")
		}

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		url := dsn + " sslmode=disable"

		fmt.Println(url)

		DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			return err
		}

		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}

		// Define 5 seconds timeout for connecting to the database
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Ping the database to ensure connection
		err = sqlDB.PingContext(ctx)
		if err != nil {
			return err
		}

		DB = DB.Debug()

		return nil
	}

	// Exponential backoff for retrying the connection
	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 60 * time.Second

	err := backoff.Retry(conn, expBackoff)
	if err != nil {
		log.Fatal(err)
	}
}
