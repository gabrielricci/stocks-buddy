package commands

import (
	"fmt"
	"log"
	"sort"

	"gabrielricci/stocks/common"
	"gabrielricci/stocks/models"
	"gabrielricci/stocks/services"

	funk "github.com/thoas/go-funk"
)

func ExecLBBMSort(env *common.Env) error {
	fmt.Println("Getting all stocks...")

        stocks := env.Repository.ListStocks()

	fmt.Println("Calculating return on capital and earnings yield, this might take a while...")

	var filteredStocks []map[string]interface{}
	filteredStocks = funk.Chain(stocks).Map(
		func(s *models.Stock) map[string]interface{} {
			return getStatements(env, s)
		}).
		Map(calculateReturnOnCapital).
		Filter(filterMinimumReturnOnCapital).
		Map(calculateEarningsYield).
		Value().([]map[string]interface{})

	fmt.Println("Sorting results...")

	sort.Slice(filteredStocks, func(i, j int) bool {
		left := filteredStocks[i]["earningsYield"].(float64)
		right := filteredStocks[j]["earningsYield"].(float64)
		return left > right
	})

	funk.Map(filteredStocks, printResults)
	return nil
}

func getStatements(env *common.Env, s *models.Stock) map[string]interface{} {
	balance := env.Repository.GetLastBalanceSheetStatement(
		s.Ticker, false,
	)

	cashflow := env.Repository.GetLastCashflowStatement(
		s.Ticker, false,
	)

	income := env.Repository.GetLastIncomeStatement(
		s.Ticker, false,
	)

	return map[string]interface{} {
		"ticker":   s.Ticker,
		"balance":  balance,
		"cashflow": cashflow,
		"income":   income,
	}
}

func calculateReturnOnCapital(
	stock map[string]interface{},
) map[string]interface{} {
	returnOnCapital := returnOnCapital(
		stock["balance"].(*models.BalanceSheetStatement),
		stock["cashflow"].(*models.CashflowStatement),
		stock["income"].(*models.IncomeStatement),
	)

	stock["returnOnCapital"] = returnOnCapital
	return stock
}

func calculateEarningsYield(
	stock map[string]interface{},
) map[string]interface{} {
	earningsYield := earningsYield(
		stock["ticker"].(string),
		stock["balance"].(*models.BalanceSheetStatement),
		stock["cashflow"].(*models.CashflowStatement),
		stock["income"].(*models.IncomeStatement),
	)

	stock["earningsYield"] = earningsYield
	return stock
}

func filterMinimumReturnOnCapital(
	stock map[string]interface{},
) bool {
	returnOnCapital := stock["returnOnCapital"].(float64)
	return (returnOnCapital > 0.15)
}

func printResults(
	stock map[string]interface{},
) map[string]interface{} {
	ticker := stock["ticker"].(string)
	returnOnCapital := stock["returnOnCapital"].(float64)
	earningsYield := stock["earningsYield"].(float64)

	fmt.Println(
		"Ticker: ", ticker,
		" | Earnings yield: ", earningsYield,
		" | Return on capital: ", returnOnCapital,
	)

	return stock
}

func netWorkingCapital(
	balance_sheet *models.BalanceSheetStatement,
	income *models.IncomeStatement,
) int64 {
	currentAssets := balance_sheet.TotalCurrentAssets
	currentLiabilities := balance_sheet.TotalCurrentLiabilities
	return (currentAssets - currentLiabilities)
}

func netFixedAssets(
	balance *models.BalanceSheetStatement,
	cashflow *models.CashflowStatement,
) int64 {
	totalAssets := balance.TotalAssets
	totalCurrentAssets := balance.TotalCurrentAssets
	intangibleAssets := balance.IntangibleAssets
	goodwill := balance.GoodWill

	return (totalAssets -
		totalCurrentAssets -
		goodwill -
		intangibleAssets)
}

func returnOnCapital(
	balance *models.BalanceSheetStatement,
	cashflow *models.CashflowStatement,
	income *models.IncomeStatement,
) float64 {
	ebit := float64(income.OperatingIncome)
	netWorkingCapital := netWorkingCapital(balance, income)
	netFixedAssets := netFixedAssets(balance, cashflow)

	tangibleCapital := float64(netWorkingCapital + netFixedAssets)

	if tangibleCapital == 0 {
		return 0.0
	}

	return ebit / tangibleCapital
}

func earningsYield(
	ticker string,
	balance *models.BalanceSheetStatement,
	cashflow *models.CashflowStatement,
	income *models.IncomeStatement,
) float64 {
	ebit := income.OperatingIncome
	enterpriseValue := enterpriseValue(
		ticker,
		balance,
		cashflow,
		income,
	)

	earningsYield := float64(ebit) / float64(enterpriseValue)
	return earningsYield
}

func enterpriseValue(
	ticker string,
	balance *models.BalanceSheetStatement,
	cashflow *models.CashflowStatement,
	income *models.IncomeStatement,
) int64 {
	quote, err := services.GetCurrentQuote(ticker)
	if err != nil {
		log.Println("ERROR: Couldn't get stock data - " + ticker)
		return 0
	}

	marketCap := quote.MarketCap
	shortDebt := balance.ShortLongTermDebt
	longDebt := balance.LongTermDebt
	minorityInterest := income.MinorityInterest
	cashAndEquivalents := balance.Cash
	shortInvestments := balance.ShortTermInvestments

	enterpriseValue := marketCap +
		shortDebt +
		longDebt +
		minorityInterest -
		cashAndEquivalents -
		shortInvestments

	return enterpriseValue
}
