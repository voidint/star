package main

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/voidint/star/plugin"
	"github.com/voidint/star/store"
)

var pullCmd = ishell.Cmd{
	Name: "pull",
	Help: "pull stars",
	Func: func(ctx *ishell.Context) {
		repos, err := plugin.PickHolder(store.CurrentHolder()).FetchAllStarredRepos()
		if err != nil {
			fmt.Println(err)
			return
		}

		store.OverwriteRepoedTag("", repos)
	},
}
