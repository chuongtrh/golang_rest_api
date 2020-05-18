package utils

import (
	"demo_api/src/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/zerolog/log"
)

func CreateConnectionDB() (*gorm.DB, error) {
	Driver := config.Cfg.DbDriver
	DbHost := config.Cfg.DbHost
	DbUser := config.Cfg.DbUser
	DbPassword := config.Cfg.DbPassword
	DbName := config.Cfg.DbName
	DbPort := config.Cfg.DbPort
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Driver, url)
	//defer db.Close()

	if err != nil {
		log.Fatal().Msgf("This is the error:", err)
	} else {
		log.Info().Msgf("We are connected to the %s database", Driver)
		db.LogMode(true)
	}

	return db, err
}
