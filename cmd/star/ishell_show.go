package main

import (
	"github.com/abiosoft/ishell"
)

var showCmd = ishell.Cmd{
	Name: "show",
	Help: "show starred repositories",
	Func: func(ctx *ishell.Context) {
		// tags := store.ListRepoedTag()

		// sb := new(strings.Builder)
		// for i := range tags {
		// 	table := tablewriter.NewWriter(sb)
		// 	table.SetHeader([]string{"Full Name", "URL"})
		// 	rows := make([][]string, 0, len(tags[i].Repos))
		// 	for j := range tags[i].Repos {
		// 		rows = append(rows, []string{tags[i].Repos[j].FullName, tags[i].Repos[j].HTMLURL})
		// 	}
		// 	table.AppendBulk(rows)
		// 	table.Render()
		// }
		// if err := ctx.ShowPaged(sb.String()); err != nil {
		// 	ctx.Println(err)
		// }
	},
}

var showGCmd = ishell.Cmd{
	Name: "G",
	Help: "display result vertically",
	Func: func(ctx *ishell.Context) {
		ctx.Println("show G haha.")
	},
}
