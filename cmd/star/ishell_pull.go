package main

import (
	"github.com/abiosoft/ishell"
)

var pullCmd = ishell.Cmd{
	Name: "pull",
	Help: "pull stars",
	Func: func(ctx *ishell.Context) {
		// 1、从远程拉取标星仓库列表
		// 2、加载节点树（若本地未加载，则从远程拉取data.json并生成节点树）
		// 3、遍历标星仓库列表，若仓库未包含在节点树中，则自动将该仓库加入节点树中的默认标签下。

		// repos, err := plugin.PickHolder(store.CurrentHolder()).FetchAllStarredRepos()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// store.OverwriteRepoedTag("", repos)
	},
}
