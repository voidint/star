package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/voidint/star/build"
	"github.com/voidint/star/plugin"
	"golang.org/x/crypto/ssh/terminal"

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
			Action: func(ctx *cli.Context) (err error) {
				var hName string
				if hName = ctx.Args().First(); hName == "" {
					hName = plugin.DefHolder
				}

				h := plugin.PickHolder(hName)
				if h == nil {
					return cli.NewExitError(fmt.Sprintf("[star] Invalid star holder name %q.", hName), 1)
				}

				auth := plugin.Authentication{
					Token: os.Getenv(fmt.Sprintf("STAR_%s_TOKEN", strings.ToUpper(hName))),
				}
				if auth.Token == "" {
					if auth.Username, auth.Password, err = readBasicAuth(); err != nil {
						return cli.NewExitError(fmt.Sprintf("[star] %q.", err.Error()), 2)
					}
				}

				user, err := h.Login(&auth)
				if err != nil {
					return cli.NewExitError(fmt.Sprintf("[star] Login failed %q.", err.Error()), 3)
				}
				runShell(hName, user)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("[star] %s", err.Error()))
		os.Exit(1)
	}
}

func readBasicAuth() (username, password string, err error) {
	state, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", "", err
	}
	defer terminal.Restore(int(os.Stdin.Fd()), state)
	term := terminal.NewTerminal(os.Stdin, "")

	term.SetPrompt("Enter Username: ")
	username, err = term.ReadLine()
	if err != nil {
		return "", "", err
	}
	password, err = term.ReadPassword("Enter Password: ")
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}
