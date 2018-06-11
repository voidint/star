package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/voidint/star/build"
	"github.com/voidint/star/plugin"

	_ "github.com/voidint/star/plugin/gitee"
	_ "github.com/voidint/star/plugin/github"
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
					hName = plugin.DefHolder
				}

				if sg := plugin.PickHolder(hName); sg == nil {
					fmt.Fprintln(os.Stderr, fmt.Sprintf("[star] Invalid star holder name %q.", hName))
					os.Exit(1)
				}

				token := os.Getenv(fmt.Sprintf("STAR_%s_TOKEN", strings.ToUpper(hName)))
				runShell(hName, token)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("[star] %s", err.Error()))
		os.Exit(1)
	}
}
