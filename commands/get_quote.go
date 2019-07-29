package commands

import (
	"fmt"

	"gabrielricci/stocks/common"
	"gabrielricci/stocks/services"
)

func ExecGetQuotesCommand(env *common.Env, ticker string) error {
	fmt.Println("Getting stock price...")

	quote, err  := services.GetCurrentQuote(ticker)
	if err != nil {
		return err
	}

	template := `
TICKER:      %s
DATE:        %s
LAST:        %f
BID:         %f
ASK:         %f
OPEN:        %f
PREV. CLOSE: %f
HIGH:        %f
LOW:         %f
`

        fmt.Printf(template,
		ticker,
		quote.Date,
		quote.Last,
		quote.Bid,
		quote.Ask,
		quote.Open,
		quote.PreviousClose,
		quote.High,
		quote.Low)

	return nil
}
