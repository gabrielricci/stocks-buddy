package services

import (
	"time"

	"gabrielricci/stocks/models"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/equity"
)

func GetCurrentQuote(ticker string) (*models.Quote, error) {
	quote, err := equity.Get(ticker)
	if err != nil {
		return nil, err
	}

	return createQuote(ticker, quote), nil
}

func createQuote(ticker string, price *finance.Equity) *models.Quote {
	return &models.Quote{
		Ticker:             ticker,
		Date:               time.Unix(int64(price.RegularMarketTime), 0),
		MarketState:        string(price.MarketState),
		Delay:              price.QuoteDelay,

		Last:               price.RegularMarketPrice,
		Bid:                price.Bid,
		Ask:                price.Ask,
		Open:               price.RegularMarketOpen,
		PreviousClose:      price.RegularMarketPreviousClose,
		High:               price.RegularMarketDayHigh,
		Low:                price.RegularMarketDayLow,

		SharesOutstanding:  price.SharesOutstanding,
		MarketCap:          price.MarketCap,
	}
}
