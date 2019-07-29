package main

import (
	"log"
	"os"

	"gabrielricci/stocks/db"
	"gabrielricci/stocks/commands"
	"gabrielricci/stocks/common"
)

func main()  {
	repo, err := db.NewSQLiteRepository("./stocks.db")
	if err != nil {
		log.Fatal("Could not connect to database")
	}
	defer repo.Close()

	env := &common.Env{Repository: repo}
	app := commands.Setup(env)

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
