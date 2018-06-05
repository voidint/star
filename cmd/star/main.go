package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/voidint/star/build"
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
			Action: func(c *cli.Context) error {
				fmt.Println("Wellcome to github star manager")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("[star] %s", err.Error()))
		os.Exit(1)
	}
}
