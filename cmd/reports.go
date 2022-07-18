package main

import (
	"github.com/urfave/cli/v2"
)

var ReportsCmd = &cli.Command{
	Name: "reports",
	Action: func(context *cli.Context) error {
		//go lib.GinServer()
		return nil
	},
}
