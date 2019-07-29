package db

import (
	"gabrielricci/stocks/models"
)

func (repo *Repository) CreateIncomeStatement(income *models.IncomeStatement) error {
	repo.Create(income)
	return nil
}

func (repo *Repository) GetLastIncomeStatement(ticker string, yearly bool) *models.IncomeStatement {
	var statement models.IncomeStatement

	repo.
		Where("ticker = ?", ticker).
		Where("yearly = ?", yearly).
		Order("period desc").
		Limit(1).
		First(&statement)

	return &statement
}
