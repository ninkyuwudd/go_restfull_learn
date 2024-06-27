package service

import (
	"database/sql"
	"gorm/learn/internal/model"
	"gorm/learn/internal/repository"
)

func GetLaptopSpec(db *sql.DB) ([]model.LaptopSpec,error){

	return repository.GetLaptopSpecRepo(db)

}