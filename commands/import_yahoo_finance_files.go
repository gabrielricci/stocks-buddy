package commands

import (
	"fmt"
	"log"
	"time"
	"strconv"
	"reflect"

	"github.com/tidwall/gjson"
	"io/ioutil"
)

const json_dir = "data/yahoo_finance/"

func ImportYahooFinanceFiles() bool {
	return execute()
}

func execute() bool {
	files, err := ioutil.ReadDir(json_dir)
	if err != nil {
		log.Println("Error reading stocks dir")
		return false
	}

	for _, file := range files {
		parseFile(json_dir + file.Name())
	}

	return true
}

func parseFile(filepath string) {
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading contents of ", filepath)
		return;
	}

	jsonData := gjson.Parse(string(contents)).Get("QuoteSummaryStore.0")
	ticker := jsonData.Get("price.symbol").String()

	fmt.Println("Parseing stock", ticker)

	var yearlyBalanceSheetsList    []*BalanceSheetStatement
	var yearlyCashflowList         []*CashflowStatement
	var yearlyIncomeList           []*IncomeStatement
	var quarterlyBalanceSheetsList []*BalanceSheetStatement
	var quarterlyCashflowList      []*CashflowStatement
	var quarterlyIncomeList        []*IncomeStatement

	yearlyBalanceSheets    := jsonData.Get("balanceSheetHistory.balanceSheetStatements")
	yearlyCashflow         := jsonData.Get("cashflowStatementHistory.cashflowStatements")
	yearlyIncome           := jsonData.Get("incomeStatementHistory.incomeStatementHistory")
	quarterlyBalanceSheets := jsonData.Get("balanceSheetHistoryQuarterly.balanceSheetStatements")
	quarterlyCashflow      := jsonData.Get("cashflowStatementHistoryQuarterly.cashflowStatements")
	quarterlyIncome        := jsonData.Get("incomeStatementHistoryQuarterly.incomeStatementHistory")

	if yearlyBalanceSheets.Exists() {
		for _, balanceSheet := range yearlyBalanceSheets.Array() {
			sheet := newBalanceSheetStatement(balanceSheet, true)
			yearlyBalanceSheetsList = append(yearlyBalanceSheetsList, sheet)
		}
	}

	if yearlyCashflow.Exists() {
		for _, cashflow := range yearlyCashflow.Array() {
			sheet := newCashflowStatement(cashflow, true)
			yearlyCashflowList = append(yearlyCashflowList, sheet)
		}
	}

	if yearlyIncome.Exists() {
		for _, income := range yearlyIncome.Array() {
			sheet := newIncomeStatement(income, true)
			yearlyIncomeList = append(yearlyIncomeList, sheet)
		}
	}

	if quarterlyBalanceSheets.Exists() {
		for _, balanceSheet := range quarterlyBalanceSheets.Array() {
			sheet := newBalanceSheetStatement(balanceSheet, false)
			quarterlyBalanceSheetsList = append(quarterlyBalanceSheetsList, sheet)
		}
	}

	if quarterlyCashflow.Exists() {
		for _, cashflow := range quarterlyCashflow.Array() {
			sheet := newCashflowStatement(cashflow, false)
			quarterlyCashflowList = append(quarterlyCashflowList, sheet)
		}
	}

	if quarterlyIncome.Exists() {
		for _, income := range quarterlyIncome.Array() {
			sheet := newIncomeStatement(income, false)
			quarterlyIncomeList = append(quarterlyIncomeList, sheet)
		}
	}

	stock := &Stock{
		Ticker: ticker,
		YearlyBalanceSheetStatements: yearlyBalanceSheetsList,
		YearlyCashflowStatements: yearlyCashflowList,
		YearlyIncomeStatements: yearlyIncomeList,
		QuarterlyBalanceSheetStatements: quarterlyBalanceSheetsList,
		QuarterlyCashflowStatements: quarterlyCashflowList,
		QuarterlyIncomeStatements: quarterlyIncomeList,
	}

	fmt.Println("Quarterly balance sheets:", len(stock.QuarterlyBalanceSheetStatements))
	fmt.Println("Quarterly cashflows:", len(stock.QuarterlyCashflowStatements))
	fmt.Println("Quarterly incomes:", len(stock.QuarterlyIncomeStatements))
	fmt.Println("Yearly balance sheets:", len(stock.YearlyBalanceSheetStatements))
	fmt.Println("Yearly cashflows:", len(stock.YearlyCashflowStatements))
	fmt.Println("Yearly incomes:", len(stock.YearlyIncomeStatements))
}

func newBalanceSheetStatement(data gjson.Result, yearly bool) *BalanceSheetStatement{
	periodRaw := time.Unix(data.Get("endDate.raw").Int(), 0)

	balanceSheetStatement := &BalanceSheetStatement{
		Period: formatPeriod(periodRaw, yearly),
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
		"LongTermInvestments": "longTermInvestments.raw",
		"TotalStockholderEquity": "totalStockholderEquity.raw",
		"Cash": "cash.raw",
		"OtherCurrentLiab": "otherCurrentLiab.raw",
		"PropertyPlantEquipment": "propertyPlantEquipment.raw",
	}, data, balanceSheetStatement)

	return balanceSheetStatement
}

func newCashflowStatement(data gjson.Result, yearly bool) *CashflowStatement{
	periodRaw := time.Unix(data.Get("endDate.raw").Int(), 0)

	cashflowStatement := &CashflowStatement{
		Period: formatPeriod(periodRaw, yearly),
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

func newIncomeStatement(data gjson.Result, yearly bool) *IncomeStatement{
	periodRaw := time.Unix(data.Get("endDate.raw").Int(), 0)

	incomeStatement := &IncomeStatement{
		Period: formatPeriod(periodRaw, yearly),
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

func formatPeriod(period time.Time, yearly bool) string {
	year := strconv.Itoa(period.Year())

	if yearly {
		return year
	} else {
		switch period.Month() {
		case 3:
			return year + ".Q1"
		case 6:
			return year + ".Q2"
		case 9:
			return year + ".Q3"
		case 12:
			return year + ".Q4"
		}
	}

	return "Unknown"
}
