package handler

import (
	"database/sql"
	"gorm/learn/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllLaptopDataHandler(db *sql.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")
		getToken := strings.Split(bearerToken, " ")[1]

		for _, token := range service.Tokens {
			if token == getToken {
				dataLaptop, _ := service.GetLaptopSpec(db)

				ctx.IndentedJSON(http.StatusOK, gin.H{
					"data": dataLaptop,
				})

				return
			}
		}

	}
}
