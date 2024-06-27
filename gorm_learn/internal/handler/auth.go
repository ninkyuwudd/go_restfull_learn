package handler

import (
	"database/sql"
	"gorm/learn/internal/model"
	"gorm/learn/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthHandler(db *sql.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var userInput model.UserInput

		if err := ctx.ShouldBindJSON(&userInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		email := userInput.Email
		pass := userInput.Password

		token, statusBool := service.Auth(db, email, pass)

		if statusBool {
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized ,please insert correct email/password"})
		}

	}
}
