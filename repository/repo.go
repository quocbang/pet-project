package repository

import (
	"fmt"

	"github.com/quocbang/pet-project/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// return client by *gorm.Client
func ConnectToDatabase(cfg config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DBHOST, 
		cfg.DBUSER, 
		cfg.DBPASSWORD, 
		cfg.DBNAME, 
		cfg.PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
