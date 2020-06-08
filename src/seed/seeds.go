package main

import (
	"demo_api/src/config"
	"demo_api/src/module/user"
	"demo_api/src/util"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func migrate(db *gorm.DB) {
	db.DropTable(&user.User{})
	db.AutoMigrate(&user.User{})
}

func seed(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	println("password:", string(password))

	var user = &user.User{
		Email:    "admin@yopmail.com",
		Password: string(password),
		Role:     user.RoleAdmin,
	}
	db.Save(user)
}

func main() {

	//Load env
	if err := config.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	if db, err := util.CreateConnectionDB(); err != nil {
	} else {
		defer db.Close()
		migrate(db)
		seed(db)
	}
}