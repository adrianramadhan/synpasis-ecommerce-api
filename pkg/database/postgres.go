package database

import (
	"fmt"
	"log"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.PostgresHost(),
		config.PostgresUser(),
		config.PostgresPassword(),
		config.PostgresDb(),
		config.PostgresPort(),
		config.PostgresSslMode(),
		config.PostgresTimezone(),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func AutoMigrate(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Cart{},
		&model.CartItem{},
		&model.Order{},
		&model.Payment{},
	)
}
