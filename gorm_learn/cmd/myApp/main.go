package main

import (
	"gorm/learn/internal/router"
	"gorm/learn/pkg/database"
)

func main() {

	db := database.InitializeDb()

	router.SetRoute(db)

}
