package service

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

var Tokens []string

func Auth(db *sql.DB, email string, pass string) (string, bool) {

	if GetAccountByEmail(db, email, pass) {
		token, _ := GenerateAuthToken()
		return token, true
	}
	return "", false

}

func GenerateAuthToken() (string, error) {
	token, err := randomHex(20)
	Tokens = append(Tokens, token)

	// gn.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })
	return Tokens[0], err
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
