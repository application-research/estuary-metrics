package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	// cli to run metrics rest daemon
	// cli to generate metrics report
	app := cli.NewApp()
	app.Description = `'estmetrics' is a cli tool to look up estuary metrics.`
	app.Name = "estmetrics"
	app.Commands = []*cli.Command{}
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug logging",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
