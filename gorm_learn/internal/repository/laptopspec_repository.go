package repository

import (
	"database/sql"
	"fmt"
	"gorm/learn/internal/model"
)

func GetLaptopSpecRepo(db *sql.DB) ([]model.LaptopSpec, error) {
	var AlllaptopSpec []model.LaptopSpec

	read, err := db.Query("select * from `laptopspec`")

	if err != nil {
		panic(fmt.Sprintf("error %s when get data from table laptopspec", err))
	}

	for read.Next() {
		var laptopSpec model.LaptopSpec
		if err = read.Scan(&laptopSpec.Id, &laptopSpec.Name, &laptopSpec.Vendor, &laptopSpec.Cpu, &laptopSpec.Ram, &laptopSpec.Storage); err != nil {
			panic(fmt.Sprintf("error %s  when scan to laptopSpec model", err))
		}
		AlllaptopSpec = append(AlllaptopSpec, laptopSpec)
	}

	return AlllaptopSpec, err

}
