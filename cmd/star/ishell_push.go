package main

import (
	"strings"

	"github.com/abiosoft/ishell"
)

var pushCmd = ishell.Cmd{
	Name: "push",
	Help: "push stars to remote",
	Func: func(ctx *ishell.Context) {
		ctx.Println("Hello", strings.Join(ctx.Args, " "))
	},
}
