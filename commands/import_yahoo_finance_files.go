package commands

import (
	"fmt"
	"time"
	"strconv"
	"reflect"
	"errors"

	"github.com/tidwall/gjson"
	"io/ioutil"

	"gabrielricci/stocks/common"
	models "gabrielricci/stocks/models"
)

func ExecImportYahooFinanceCommand(env *common.Env, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return errors.New("Could not read stocks path")
	}

	for _, file := range files {
		parseFile(env, path + file.Name())
	}

	return nil
}

func parseFile(env *common.Env, filepath string) {
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ERROR: Could not read  contents of", filepath)
		return;
	}

	jsonData := gjson.Parse(string(contents)).Get("QuoteSummaryStore.0")
	ticker := jsonData.Get("price.symbol").String()
	fmt.Println("Importing", ticker)


	stock := &models.Stock{
		Ticker: ticker,
	}

	fmt.Println(stock.Ticker)

	err = env.Repository.CreateStock(stock)
	if err != nil {
		fmt.Println("Error saving the stock data:", err)
		return;
	}

	var balanceSheetsList    []*models.BalanceSheetStatement
	var cashflowList         []*models.CashflowStatement
	var incomeList           []*models.IncomeStatement

	yearlyBalanceSheets    := jsonData.Get("balanceSheetHistory.balanceSheetStatements")
	yearlyCashflow         := jsonData.Get("cashflowStatementHistory.cashflowStatements")
	yearlyIncome           := jsonData.Get("incomeStatementHistory.incomeStatementHistory")
	quarterlyBalanceSheets := jsonData.Get("balanceSheetHistoryQuarterly.balanceSheetStatements")
	quarterlyCashflow      := jsonData.Get("cashflowStatementHistoryQuarterly.cashflowStatements")
	quarterlyIncome        := jsonData.Get("incomeStatementHistoryQuarterly.incomeStatementHistory")

	if yearlyBalanceSheets.Exists() {
		for _, balanceSheet := range yearlyBalanceSheets.Array() {
			statement := newBalanceSheetStatement(ticker, balanceSheet, true)
			env.Repository.CreateBalanceSheetStatement(statement)
			balanceSheetsList = append(balanceSheetsList, statement)
		}
	}

	if yearlyCashflow.Exists() {
		for _, cashflow := range yearlyCashflow.Array() {
			statement := newCashflowStatement(ticker, cashflow, true)
			env.Repository.CreateCashflowStatement(statement)
			cashflowList = append(cashflowList, statement)
		}
	}

	if yearlyIncome.Exists() {
		for _, income := range yearlyIncome.Array() {
			statement := newIncomeStatement(ticker, income, true)
			env.Repository.CreateIncomeStatement(statement)
			incomeList = append(incomeList, statement)
		}
	}

	if quarterlyBalanceSheets.Exists() {
		for _, balanceSheet := range quarterlyBalanceSheets.Array() {
			statement := newBalanceSheetStatement(ticker, balanceSheet, false)
			env.Repository.CreateBalanceSheetStatement(statement)
			balanceSheetsList = append(balanceSheetsList, statement)
		}
	}

	if quarterlyCashflow.Exists() {
		for _, cashflow := range quarterlyCashflow.Array() {
			statement := newCashflowStatement(ticker, cashflow, false)
			env.Repository.CreateCashflowStatement(statement)
			cashflowList = append(cashflowList, statement)
		}
	}

	if quarterlyIncome.Exists() {
		for _, income := range quarterlyIncome.Array() {
			statement := newIncomeStatement(ticker, income, false)
			env.Repository.CreateIncomeStatement(statement)
			incomeList = append(incomeList, statement)
		}
	}

}

func newBalanceSheetStatement(ticker string, data gjson.Result, yearly bool) *models.BalanceSheetStatement{
	period := time.Unix(data.Get("endDate.raw").Int(), 0)

	balanceSheetStatement := &models.BalanceSheetStatement{
		Ticker: ticker,
		Period: &period,
		Yearly: yearly,
	}

	fillNumericValues(map[string]string{
		"LongTermDebt": "longTermDebt.raw",
		"TotalLiab": "totalLiab.raw",
		"TotalCurrentAssets": "totalCurrentAssets.raw",
		"TotalAssets": "totalAssets.raw",
		"OtherStockholderEquity": "otherStockholderEquity.raw",
		"ShortLongTermDebt": "shortLongTermDebt.raw",
		"NetReceivables": "netReceivables.raw",
		"TotalCurrentLiabilities": "totalCurrentLiabilities.raw",
		"OtherCurrentAssets": "otherCurrentAssets.raw",
		"NetTangibleAssets": "netTangibleAssets.raw",
		"TreasuryStock": "treasuryStock.raw",
		"OtherLiab": "otherLiab.raw",
		"Inventory": "inventory.raw",
		"CapitalSurplus": "capitalSurplus.raw",
		"OtherAssets": "otherAssets.raw",
		"AccountsPayable": "accountsPayable.raw",
		"GoodWill": "goodWill.raw",
		"IntangibleAssets": "intangibleAssets.raw",
		"CommonStock": "commonStock.raw",
		"DeferredLongTermAssetCharges": "deferredLongTermAssetCharges.raw",
		"RetainedEarnings": "retainedEarnings.raw",
		"ShortTermInvestments": "shortTermInvestments.raw",
		"LongTermInvestments": "longTermInvestments.raw",
		"TotalStockholderEquity": "totalStockholderEquity.raw",
		"Cash": "cash.raw",
		"OtherCurrentLiab": "otherCurrentLiab.raw",
		"PropertyPlantEquipment": "propertyPlantEquipment.raw",
	}, data, balanceSheetStatement)

	return balanceSheetStatement
}

func newCashflowStatement(ticker string, data gjson.Result, yearly bool) *models.CashflowStatement{
	period := time.Unix(data.Get("endDate.raw").Int(), 0)

	cashflowStatement := &models.CashflowStatement{
		Ticker: ticker,
		Period: &period,
		Yearly: yearly,
	}

	fillNumericValues(map[string]string{
		"CapitalExpenditures": "capitalExpenditures.raw",
		"ChangeToAccountReceivables": "changeToAccountReceivables.raw",
		"NetIncome": "netIncome.raw",
		"ChangeToNetincome": "changeToNetincome.raw",
		"ChangeToLiabilities": "changeToLiabilities.raw",
		"TotalCashFromFinancingActivities": "totalCashFromFinancingActivities.raw",
		"NetBorrowings": "netBorrowings.raw",
		"Depreciation": "depreciation.raw",
		"ChangeInCash": "changeInCash.raw",
		"TotalCashflowsFromInvestingActivities": "totalCashflowsFromInvestingActivities.raw",
		"TotalCashFromOperatingActivities": "totalCashFromOperatingActivities.raw",
		"ChangeToOperatingActivities": "changeToOperatingActivities.raw",
		"OtherCashflowsFromFinancingActivities": "otherCashflowsFromFinancingActivities.raw",
		"DividendsPaid": "dividendsPaid.raw",
	}, data, cashflowStatement)

	return cashflowStatement
}

func newIncomeStatement(ticker string, data gjson.Result, yearly bool) *models.IncomeStatement{
	period := time.Unix(data.Get("endDate.raw").Int(), 0)

	incomeStatement := &models.IncomeStatement{
		Ticker: ticker,
		Period: &period,
		Yearly: yearly,
	}

	fillNumericValues(map[string]string{
		"IncomeTaxExpense": "incomeTaxExpense.raw",
		"TotalOperatingExpenses": "totalOperatingExpenses.raw",
		"ResearchDevelopment": "researchDevelopment.raw",
		"TotalOtherIncomeExpenseNet": "totalOtherIncomeExpenseNet.raw",
		"NetIncomeApplicableToCommonShares": "netIncomeApplicableToCommonShares.raw",
		"OtherOperatingExpenses": "otherOperatingExpenses.raw",
		"OtherItems": "otherItems.raw",
		"GrossProfit": "grossProfit.raw",
		"EffectOfAccountingCharges": "effectOfAccountingCharges.raw",
		"ExtraordinaryItems": "extraordinaryItems.raw",
		"IncomeBeforeTax": "incomeBeforeTax.raw",
		"TotalRevenue": "totalRevenue.raw",
		"NetIncome": "netIncome.raw",
		"MaxAge": "maxAge.raw",
		"Ebit": "ebit.raw",
		"MinorityInterest": "minorityInterest.raw",
		"NonRecurring": "nonRecurring.raw",
		"DiscontinuedOperations": "discontinuedOperations.raw",
		"InterestExpense": "interestExpense.raw",
		"CostOfRevenue": "costOfRevenue.raw",
		"SellingGeneralAdministrative": "sellingGeneralAdministrative.raw",
		"OperatingIncome": "operatingIncome.raw",
		"NetIncomeFromContinuingOps": "netIncomeFromContinuingOps.raw",
	}, data, incomeStatement)

	return incomeStatement
}

func fillNumericValues(mapping map[string]string, data gjson.Result, structure interface{}) {
	t := reflect.ValueOf(structure).Elem()
	for k, v := range mapping {
		value := data.Get(v).Int()
		t.FieldByName(k).Set(reflect.ValueOf(value))
	}
}

func isLastDayOfMonth(period *time.Time) bool {
	firstOfMonth := time.Date(period.Year(), period.Month(), 1, 0, 0, 0, 0, period.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return int(period.Day()) == int(lastOfMonth.Day())
}

func isLastMonthOfQuarter(period *time.Time) bool {
	return period.Month() == 3 || period.Month() == 6 || period.Month() == 9 || period.Month() == 12
}

func formatPeriod(period time.Time, yearly bool) string {
	year := strconv.Itoa(period.Year())

	if yearly {
		return year
	}

	if isLastDayOfMonth(&period) && isLastMonthOfQuarter(&period) {
		return year + ".Q" + strconv.Itoa(int(period.Month()) / 3)
	}


	switch period.Month() {
	case 1, 2, 3:
		lastYear := strconv.Itoa(period.Year() - 1)
		return lastYear  + ".Q4"
	case 4, 5, 6:
		return year + ".Q1"
	case 7, 8, 9:
		return year + ".Q2"
	case 10, 11, 12:
		return year + ".Q3"
	}

	return "Unknown"
}
