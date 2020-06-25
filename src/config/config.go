package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Env             string
	AppPort         string
	DbHost          string
	DbDriver        string
	DbUser          string
	DbPassword      string
	DbName          string
	DbPort          string
	JwtKey          string
	JwtExp          int
	LogFileLocation string
}

var Cfg = Config{}

func Load() error {
	err := godotenv.Load()
	if err == nil {
		Cfg = Config{
			Env:             os.Getenv("ENV"),
			AppPort:         os.Getenv("APP_PORT"),
			DbHost:          os.Getenv("DB_HOST"),
			DbDriver:        os.Getenv("DB_DRIVER"),
			DbUser:          os.Getenv("DB_USER"),
			DbPassword:      os.Getenv("DB_PASSWORD"),
			DbName:          os.Getenv("DB_NAME"),
			DbPort:          os.Getenv("DB_PORT"),
			JwtKey:          os.Getenv("API_JWT_KEY"),
			JwtExp:          cast.ToInt(os.Getenv("API_JWT_EXP")),
			LogFileLocation: os.Getenv("LOG_FILE_LOCATION"),
		}
	}
	return err
}
