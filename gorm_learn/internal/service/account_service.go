package service

import (
	"database/sql"
	"gorm/learn/internal/repository"

	// "net/http"

	// "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAccountByEmail(db *sql.DB, email string, pass string) bool {

	dataAccount, _ := repository.GetAccount(db, email)

	err := bcrypt.CompareHashAndPassword([]byte(dataAccount.Password), []byte(pass))

	return err == nil

	// if err != nil {
	// d.IndentedJSON(http.StatusUnauthorized, gin.H{
	// 	"message": "Unauthorized wrong email/pass ",
	// })
	// return false
	// panic(fmt.Sprintf("Error %s,occure when select account", err))
	// }
	// return true
}
