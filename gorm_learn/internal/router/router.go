package router

import (
	"database/sql"
	"gorm/learn/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetRoute(db *sql.DB) {

	router := gin.Default()

	router.POST("/auth", handler.AuthHandler(db))
	router.GET("/laptopSpec", handler.GetAllLaptopDataHandler(db))
	// router.POST("/register", registerPersonData)
	// router.GET("/dataTest", getDataTest)
	router.Run("localhost:8080")
}
