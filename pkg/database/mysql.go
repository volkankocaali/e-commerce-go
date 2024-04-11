package database

import (
	"fmt"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewMysqlDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%t", cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDatabase, true)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	migrate := db.AutoMigrate(
		models.Users{},
		models.Address{},
		models.Categories{},
		models.Product{},
		models.ProductCategories{},
		models.ProductTags{},
		models.Tags{},
		models.Wallet{},
		models.SlideShows{},
	)

	if err != nil {
		log.Fatalf("Error migrating database: %v", migrate.Error)
	}
	return db, err
}
