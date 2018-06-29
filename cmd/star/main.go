package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/voidint/star/plugin/gitee"
	_ "github.com/voidint/star/plugin/github"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
	"github.com/voidint/star/build"
	"github.com/voidint/star/plugin"
	"github.com/voidint/star/store"
	"golang.org/x/crypto/ssh/terminal"
)

const shortVersion = "0.1.0"

var (
	flog *os.File
)

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

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-file",
			Value: filepath.Join(os.TempDir(), "star.log"),
			Usage: "log file path",
		},
	}

	app.Before = func(ctx *cli.Context) (err error) {
		flog, err = os.Create(ctx.GlobalString("log-file"))
		if err != nil {
			return err
		}

		log.Logger = log.Output(flog)
		return nil
	}

	app.After = func(ctx *cli.Context) error {
		if flog != nil {
			return flog.Close()
		}
		return nil
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

				store.Use(hName)

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
