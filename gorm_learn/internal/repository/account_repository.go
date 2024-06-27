package repository

import (
	"database/sql"
	"fmt"
	"gorm/learn/internal/model"
)

func GetAccount(db *sql.DB, email string) (model.Account, error) {
	var account model.Account
	read, err := db.Query(fmt.Sprintf("SELECT * FROM `useraccount` WHERE `email` = '%s'", email))

	if err != nil {
		panic(fmt.Sprintf("Error %s,occure when select account", err))
	}

	for read.Next() {
		err = read.Scan(&account.Id, &account.Name, &account.Email, &account.Password)
	}

	if err != nil {
		panic(fmt.Sprintf("Error %s,occure when scan to account model", err))
	}

	return account, err
}
