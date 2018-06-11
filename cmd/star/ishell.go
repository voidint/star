package main

import (
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/voidint/star/plugin"
)

func runShell(holderName, token string) {
	shell := ishell.New()

	shell.Printf("Welcome To %s Star Interactive Shell\n", strings.Title(holderName))

	holder := plugin.PickHolder(holderName)
	holder.SetToken(token)
	plugin.Use(holderName)

	shell.AddCmd(&initCmd)
	shell.AddCmd(&pullCmd)
	shell.AddCmd(&pushCmd)

	shell.Run()
}
