package main

import (
	"fmt"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
	"github.com/voidint/star/plugin"
)

func runShell(holderName, token string) {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: fmt.Sprintf("star@%s> ", holderName),
	})

	shell.Printf("Welcome To %s Star Interactive Shell\n", strings.Title(holderName))

	holder := plugin.PickHolder(holderName)
	holder.SetToken(token)
	// plugin.Use(holderName)

	shell.AddCmd(&initCmd)
	shell.AddCmd(&pullCmd)
	shell.AddCmd(&pushCmd)

	shell.Run()
}
