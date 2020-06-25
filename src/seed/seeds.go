package main

import (
	"demo_api/src/config"
	"demo_api/src/module/user"
	"demo_api/src/util"
	"demo_api/src/util/logger"
	"github.com/jinzhu/gorm"
)

func migrate(db *gorm.DB) {
	db.DropTable(&user.User{})
	db.AutoMigrate(&user.User{})
}

func seed(db *gorm.DB) {
	password := util.HashPassword("12345678")

	var userData = &user.User{
		Email:    "admin@yopmail.com",
		Password: password,
		Role:     user.RoleAdmin,
	}
	db.Save(userData)
}

func main() {

	//Load env
	if err := config.Load(); err != nil {
		logger.Fatalf("Error getting env, %v", err)
	} else {
		logger.Info("We are getting the env values")
	}
	if db, err := util.CreateConnectionDB(); err != nil {
	} else {
		defer db.Close()
		migrate(db)
		seed(db)
	}
}
