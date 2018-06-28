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
		h := plugin.PickHolder(store.CurrentHolder())

		if _, err := h.GetRepo(); err == plugin.ErrRepoNotFound {
			if _, err = h.CreateRepo(); err != nil {
				ctx.Println(err)
				return
			}
		}

		repos, err := h.FetchAllStarredRepos()
		if err != nil {
			ctx.Println(err)
			return
		}
		store.OverwriteRepoedTag(tag.Default().Name, repos)

		repoedTags := store.ListRepoedTag()
		readme, _ := view.GenReadme(repoedTags)
		data, _ := view.GenJSONData(repoedTags)

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
		if _, err = h.UpdateFile(readmePath, fReadme.SHA, readme); err != nil {
			ctx.Println(err)
			return
		}

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
		if _, err = h.UpdateFile(dataPath, fData.SHA, data); err != nil {
			ctx.Println(err)
			return
		}

		// 1、判断仓库是否存在。若不存在则创建。
		// 2、准备data.json数据。
		// 3、依据data.json生成README.md内容。
		// 4、判断README.md是否存在。若不存在则新增，反之则更新。
		// 5、判断data.json文件是否存在。若不存在则新增，反之则更新。

	},
}
