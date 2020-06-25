package util

import (
	"demo_api/src/config"
	"demo_api/src/util/logger"
	"fmt"

	"github.com/jinzhu/gorm"
)

// CreateConnectionDB func
func CreateConnectionDB() (*gorm.DB, error) {
	Driver := config.Cfg.DbDriver
	DbHost := config.Cfg.DbHost
	DbUser := config.Cfg.DbUser
	DbPassword := config.Cfg.DbPassword
	DbName := config.Cfg.DbName
	DbPort := config.Cfg.DbPort
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Driver, url)
	// defer db.Close()

	logger.Infof("DbHost %s", DbHost)

	if err != nil {
		logger.Panicf("This is the error: %s", err)
	} else {
		logger.Infof("We are connected to the %s database", Driver)
		db.LogMode(true)
	}

	return db, err
}
