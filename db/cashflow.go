package db

import (
	"gabrielricci/stocks/models"
)

func (repo *Repository) CreateCashflowStatement(cashflow *models.CashflowStatement) error {
	repo.Create(cashflow)
	return nil
}

func (repo *Repository) GetLastCashflowStatement(ticker string, yearly bool) *models.CashflowStatement {
	var statement models.CashflowStatement

	repo.
		Where("ticker = ?", ticker).
		Where("yearly = ?", yearly).
		Order("period desc").
		Limit(1).
		First(&statement)

	return &statement
}
