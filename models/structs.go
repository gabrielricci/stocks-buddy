package models

import(
	"time"

	"github.com/jinzhu/gorm"
)

type Quote struct {
	Ticker             string
	Date               time.Time
	MarketState        string
	Delay              int

	Last               float64
	Bid                float64
	Ask                float64
	Open               float64
	PreviousClose      float64
	High               float64
	Low                float64

	SharesOutstanding  int
	MarketCap          int64
}

type HistoricalQuote struct {
	Ticker    string
	Date      time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
	AdjClose  float64
}


type Stock struct {
	gorm.Model

	Ticker                 string `gorm:"primary_key"`

	BalanceSheetStatements []*BalanceSheetStatement
	CashflowStatements     []*CashflowStatement
	IncomeStatements       []*IncomeStatement
}

type BalanceSheetStatement struct {
	gorm.Model
	Ticker string `gorm:"unique_index:uq_balance_ticker_period"`
	Period *time.Time `gorm:"unique_index:uq_balance_ticker_period"`
	Yearly bool `gorm:"unique_index:uq_balance_ticker_period"`
	LongTermDebt int64
	TotalLiab int64
	TotalCurrentAssets int64
	TotalAssets int64
	OtherStockholderEquity int64
	ShortLongTermDebt int64
	NetReceivables int64
	TotalCurrentLiabilities int64
	OtherCurrentAssets int64
	NetTangibleAssets int64
	TreasuryStock int64
	OtherLiab int64
	Inventory int64
	CapitalSurplus int64
	OtherAssets int64
	AccountsPayable int64
	GoodWill int64
	IntangibleAssets int64
	CommonStock int64
	DeferredLongTermAssetCharges int64
	RetainedEarnings int64
	ShortTermInvestments int64
	LongTermInvestments int64
	TotalStockholderEquity int64
	Cash int64
	OtherCurrentLiab int64
	PropertyPlantEquipment int64
}

type CashflowStatement struct {
	gorm.Model
	Ticker string `gorm:"unique_index:uq_cashflow_ticker_period"`
	Period *time.Time `gorm:"unique_index:uq_cashflow_ticker_period"`
	Yearly bool `gorm:"unique_index:uq_cashflow_ticker_period"`
	CapitalExpenditures int64
	ChangeToAccountReceivables int64
	NetIncome int64
	ChangeToNetincome int64
	ChangeToLiabilities int64
	TotalCashFromFinancingActivities int64
	NetBorrowings int64
	Depreciation int64
	ChangeInCash int64
	TotalCashflowsFromInvestingActivities int64
	TotalCashFromOperatingActivities int64
	ChangeToOperatingActivities int64
	OtherCashflowsFromFinancingActivities int64
	DividendsPaid int64
}

type IncomeStatement struct {
	gorm.Model
	Ticker string `gorm:"unique_index:uq_income_ticker_period"`
	Period *time.Time `gorm:"unique_index:uq_income_ticker_period"`
	Yearly bool `gorm:"unique_index:uq_income_ticker_period"`
	IncomeTaxExpense int64
	TotalOperatingExpenses int64
	ResearchDevelopment int64
	TotalOtherIncomeExpenseNet int64
	NetIncomeApplicableToCommonShares int64
	OtherOperatingExpenses int64
	OtherItems int64
	GrossProfit int64
	EffectOfAccountingCharges int64
	ExtraordinaryItems int64
	IncomeBeforeTax int64
	TotalRevenue int64
	NetIncome int64
	MaxAge int64
	Ebit int64
	MinorityInterest int64
	NonRecurring int64
	DiscontinuedOperations int64
	InterestExpense int64
	CostOfRevenue int64
	SellingGeneralAdministrative int64
	OperatingIncome int64

	NetIncomeFromContinuingOps int64
}
