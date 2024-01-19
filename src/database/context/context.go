package context

import (
	"fmt"

	Models "first_api/src/database/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var dataBase *gorm.DB

func ConntectDataBase() {
	dns := "sqlserver://amir:********@localhost:1433?database=Go_DB"
	db, err := gorm.Open(sqlserver.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Faild to connect to sqlserver")
	} else {
		fmt.Println("Succesfully Connect to sqlserver")
	}

	dataBase = db

	fmt.Println(MigrateTables(db))
}

func MigrateTables(db *gorm.DB) string {
	err := db.AutoMigrate(&Models.Product{})
	if err != nil {
		return "Failed to migrate"
	}
	return "Succesfully to migrate tables"
}

func DataBase() *gorm.DB {
	return dataBase
}
