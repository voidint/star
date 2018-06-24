package main

import (
	"github.com/abiosoft/ishell"
)

var initCmd = ishell.Cmd{
	Name: "init",
	Help: "pull stars then create repo",
	Func: func(ctx *ishell.Context) {
		// h := plugin.PickHolder(store.CurrentHolder())

		// repos, err := h.FetchAllStarredRepos()
		// if err != nil {
		// 	ctx.Println(err)
		// 	return
		// }

		// 1、判断仓库是否存在。若不存在则创建。
		// 2、准备data.json数据。
		// 3、依据data.json生成README.md内容。
		// 4、判断README.md是否存在。若不存在则新增，反之则更新。
		// 5、判断data.json文件是否存在。若不存在则新增，反之则更新。

	},
}
