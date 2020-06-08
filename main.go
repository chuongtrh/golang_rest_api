package main

import (
	app "demo_api/src"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	app.Run()
}
