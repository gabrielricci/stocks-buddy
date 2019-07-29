package db

import (
	"gabrielricci/stocks/models"
)

func (repo *Repository) CreateStock(stock *models.Stock) error {
	repo.Where(models.Stock{Ticker: stock.Ticker}).FirstOrCreate(stock)
	return nil
}

func (repo *Repository) ListStocks() []*models.Stock {
	var stocks []*models.Stock
	repo.Find(&stocks)
	return stocks
}
