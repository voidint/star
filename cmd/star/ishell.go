package main

import (
	"strings"

	"github.com/abiosoft/ishell"
)

func runShell(plugin string) {
	shell := ishell.New()

	shell.Printf("Welcome To %s Star Interactive Shell\n", strings.Title(plugin))

	shell.AddCmd(&initCmd)
	shell.AddCmd(&pullCmd)
	shell.AddCmd(&pushCmd)

	shell.Run()
}
