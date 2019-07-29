package db

import (
	"log"

	"gabrielricci/stocks/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewSQLiteRepository(dataSourceName string) (*Repository, error) {
	db, err := gorm.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Stock{})
	db.AutoMigrate(&models.BalanceSheetStatement{})
	db.AutoMigrate(&models.CashflowStatement{})
	db.AutoMigrate(&models.IncomeStatement{})

	log.Println("Connected to database")

	return &Repository{db}, nil
}
