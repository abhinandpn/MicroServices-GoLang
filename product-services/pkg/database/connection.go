package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/config"
	"github.com/abhinandpn/MicroServices-GoLang/product-services/pkg/models"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		return nil, err
	}
	return db, dbErr
}
