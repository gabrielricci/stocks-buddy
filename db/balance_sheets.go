package db

import (
	"gabrielricci/stocks/models"
)

func (repo *Repository) CreateBalanceSheetStatement(balance *models.BalanceSheetStatement) error {
	repo.Create(balance)
	return nil
}

func (repo *Repository) GetLastBalanceSheetStatement(ticker string, yearly bool) *models.BalanceSheetStatement {
	var statement models.BalanceSheetStatement

	repo.
		Where("ticker = ?", ticker).
		Where("yearly = ?", yearly).
		Order("period desc").
		Limit(1).
		First(&statement)

	return &statement
}
