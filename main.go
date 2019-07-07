package main

import (
	"fmt"
	"gabrielricci/stocks/db"
	"gabrielricci/stocks/commands"
)

func main()  {
	fmt.Println("Gabriel Ricci's stock control")
	db.OpenConnection()
	commands.ImportYahooFinanceFiles()

}
