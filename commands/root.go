package commands

import (
	"errors"

	"gabrielricci/stocks/common"

	"github.com/urfave/cli"
)



func Setup(env *common.Env) *cli.App {
	importYahooFinanceCmd := cli.Command{
		Name:    "import-yahoo-finance",
		Aliases: []string{"import-yahoo"},
		Flags:   []cli.Flag {
			cli.StringFlag{
				Name:  "path, p",
				Usage: "Path of the files to import",
			},
		},
		Usage:   "Import yahoo finance files",
		Action:  func(c *cli.Context) error {
			if c.String("path") == "" {
				return errors.New("--path is required")
			}

			return ExecImportYahooFinanceCommand(env, c.String("path"))
		},
	}

	currentQuoteCommand := cli.Command{
		Name:      "current-quote",
		Aliases:   []string{"quote", "q"},
		Usage:     "Show current quote of a `stock`",
		ArgsUsage: "[stock]",
		Action:    func (c *cli.Context) error {
			if c.NArg() < 1 {
				return errors.New("Stock not specified")
			}

			return ExecGetQuotesCommand(env, c.Args().Get(0))
		},
	}

	lbbmSortCommand := cli.Command{
		Name:    "lbbm-sort",
		Aliases: []string{"lbbm"},
		Usage:   "Sort stocks based on the `The little book that beats the markets` algorithm",
		Action:  func(c *cli.Context) error {
			return ExecLBBMSort(env)
		},
	}

	app := cli.NewApp()
	app.Name = "stocks-buddy"
	app.Usage = "A CLI helper to beat the stock market"
	app.Commands = []cli.Command{
		importYahooFinanceCmd,
		currentQuoteCommand,
		lbbmSortCommand,
	}

	return app
}
