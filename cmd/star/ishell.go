package main

import (
	"fmt"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
	"github.com/voidint/star/plugin"
)

func runShell(holderName string, user *plugin.User) {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: fmt.Sprintf("[%s@%s] $ ", user.Login, holderName),
	})

	shell.Printf("Welcome To %s Star Interactive Shell\n", strings.Title(holderName))

	shell.AddCmd(&initCmd)
	shell.AddCmd(&pullCmd)
	shell.AddCmd(&pushCmd)

	shell.Run()
}
