package main

import (
	"strings"

	"github.com/abiosoft/ishell"
)

var initCmd = ishell.Cmd{
	Name: "init",
	Help: "pull stars then create repo",
	Func: func(ctx *ishell.Context) {
		ctx.Println("Hello", strings.Join(ctx.Args, " "))
	},
}
