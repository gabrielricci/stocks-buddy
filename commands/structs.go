package commands

type Stock struct {
	Ticker                          string

	YearlyBalanceSheetStatements    []*BalanceSheetStatement
	YearlyCashflowStatements        []*CashflowStatement
	YearlyIncomeStatements          []*IncomeStatement

	QuarterlyBalanceSheetStatements []*BalanceSheetStatement
	QuarterlyCashflowStatements     []*CashflowStatement
	QuarterlyIncomeStatements       []*IncomeStatement
}

type BalanceSheetStatement struct {
	Period string
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
	LongTermInvestments int64
	TotalStockholderEquity int64
	Cash int64
	OtherCurrentLiab int64
	PropertyPlantEquipment int64
}

type CashflowStatement struct {
	Period string
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
	Period string
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
