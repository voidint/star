package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/voidint/star/build"
	"github.com/voidint/star/holder"

	_ "github.com/voidint/star/holder/gitee"
	_ "github.com/voidint/star/holder/github"
)

const shortVersion = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "star"
	app.Usage = "Github star manager."
	app.Version = build.Version(shortVersion)
	app.Copyright = "Copyright (c) 2018, 2018, voidint. All rights reserved."
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "voidnt",
			Email: "voidint@126.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Login the interactive environment",
			Action: func(ctx *cli.Context) error {
				var hName string
				if hName = ctx.Args().First(); hName == "" {
					hName = holder.DefHolder
				}

				if sg := holder.PickStargazer(hName); sg == nil {
					fmt.Fprintln(os.Stderr, fmt.Sprintf("[star] Invalid star holder name %q.", hName))
					os.Exit(1)
				}

				runShell(hName)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("[star] %s", err.Error()))
		os.Exit(1)
	}
}
