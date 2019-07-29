package db

import (
	"github.com/jinzhu/gorm"
	models "gabrielricci/stocks/models"
)

type Datastore interface {
	CreateStock(*models.Stock) error
	CreateBalanceSheetStatement(*models.BalanceSheetStatement) error
	CreateCashflowStatement(*models.CashflowStatement) error
	CreateIncomeStatement(*models.IncomeStatement) error

	GetLastBalanceSheetStatement(string, bool) *models.BalanceSheetStatement
	GetLastCashflowStatement(string, bool) *models.CashflowStatement
	GetLastIncomeStatement(string, bool) *models.IncomeStatement

	ListStocks() []*models.Stock
}

type Repository struct {
	*gorm.DB
}
