package main

import (
	"strings"

	"github.com/abiosoft/ishell"
)

var pullCmd = ishell.Cmd{
	Name: "pull",
	Help: "pull stars",
	Func: func(ctx *ishell.Context) {
		ctx.Println("Hello", strings.Join(ctx.Args, " "))
	},
}
