package main

import (
	"github.com/abiosoft/ishell"
	"github.com/voidint/star/plugin"
	"github.com/voidint/star/store"
	"github.com/voidint/star/tag"
	"github.com/voidint/star/view"
)

const (
	readmePath = "README.md"
	dataPath   = "data.json"
)

var initCmd = ishell.Cmd{
	Name: "init",
	Help: "pull stars then create repo",
	Func: func(ctx *ishell.Context) {
		// 0、获取当前holder（github/gitee）
		h := plugin.PickHolder(store.CurrentHolder())

		// 1、判断仓库是否存在，若不存在则创建。
		if _, err := h.GetRepo(); err == plugin.ErrRepoNotFound {
			if _, err = h.CreateRepo(); err != nil {
				ctx.Println(err)
				return
			}
		}

		// 2、从远程拉取所欲标星仓库并存入本地默认标签下
		repos, err := h.FetchAllStarredRepos()
		if err != nil {
			ctx.Println(err)
			return
		}
		// TODO 清空除默认节点(标签+标星仓库)外的所有节点
		store.OverwriteRepos(tag.DefaultPath, repos)

		// 3、读取本地所有节点(标签+标星仓库)并生成README.md与data.json文件内容
		nodes := store.ListNodes()
		readme, _ := view.GenReadme(nodes)
		data, _ := view.GenJSONData(nodes)

		// 4、判断远程仓库中是否含有README.md，若不存在则创建。
		fReadme, err := h.GetFile(readmePath)
		if err == plugin.ErrFileNotFound {
			if fReadme, err = h.CreateFile(readmePath, readme); err != nil {
				ctx.Println(err)
				return
			}
		}
		if err != nil {
			ctx.Println(err)
			return
		}
		// 5、更新远程的README.md文件内容
		if _, err = h.UpdateFile(readmePath, fReadme.SHA, readme); err != nil {
			ctx.Println(err)
			return
		}

		// 6、判断远程仓库中是否含有data.json，若不存在则创建。
		fData, err := h.GetFile(dataPath)
		if err == plugin.ErrFileNotFound {
			if fData, err = h.CreateFile(dataPath, data); err != nil {
				ctx.Println(err)
				return
			}
		}
		if err != nil {
			ctx.Println(err)
			return
		}
		// 7、更新远程的data.json文件内容
		if _, err = h.UpdateFile(dataPath, fData.SHA, data); err != nil {
			ctx.Println(err)
			return
		}
	},
}
