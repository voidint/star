package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/rs/zerolog/log"
	"github.com/voidint/star/plugin"
	"github.com/voidint/star/store"
	"github.com/voidint/star/tag"
	"github.com/voidint/star/view"
)

var pullCmd = ishell.Cmd{
	Name: "pull",
	Help: "pull stars",
	Func: func(ctx *ishell.Context) {
		ctx.ProgressBar().Start()

		// 0、获取当前holder（github/gitee）
		h := plugin.PickHolder(store.CurrentHolder())
		ctx.ProgressBar().Suffix(fmt.Sprint(" ", 10, "%"))
		ctx.ProgressBar().Progress(10)

		// 1、从远程仓库拉取data.json数据
		fData, err := h.GetFile(dataPath)
		ctx.ProgressBar().Suffix(fmt.Sprint(" ", 30, "%"))
		ctx.ProgressBar().Progress(30)
		if err == plugin.ErrFileNotFound {
			// 2、若仓库不存在或者data.json不存在，则终止流程并提示先执行init命令。
			log.Warn().Msgf("Fetch data.json error: %s", err.Error())
			ctx.Println("Execute the `init` command first.")
			return
		}
		// 3、从远程拉取标星仓库列表
		repos, err := h.FetchAllStarredRepos()
		ctx.ProgressBar().Suffix(fmt.Sprint(" ", 60, "%"))
		ctx.ProgressBar().Progress(60)
		if err != nil {
			log.Error().Msgf("Fetch starred repos error: %s", err.Error())
			ctx.Println(err)
			return
		}

		// 4、依据远程的data.json和标星仓库列表生成远程节点树。
		jsonData, err := base64.StdEncoding.DecodeString(fData.Content)
		var rNodes []*store.Node
		if err = json.Unmarshal(jsonData, &rNodes); err != nil {
			log.Error().Str("data.json", string(jsonData)).Msgf("Unmarshal data.json error: %s", err.Error())
			ctx.Println(err)
			return
		}
		ctx.ProgressBar().Suffix(fmt.Sprint(" ", 70, "%"))
		ctx.ProgressBar().Progress(70)

		var newrepos []*plugin.Repo
		for i := range repos {
			if !strings.Contains(string(jsonData), repos[i].FullName) { // TODO 不够严谨
				log.Debug().Msgf("Untagged repo: %s", repos[i].FullName)
				newrepos = append(newrepos, repos[i])
				continue
			}
			log.Debug().Msgf("Tagged repo: %s", repos[i].FullName)
		}
		if len(newrepos) > 0 {
			defNode := store.Nodes(rNodes).FindNode(tag.DefaultPath)
			if defNode == nil {
				defNode = &store.Node{
					Tag: &tag.DefaultTag,
				}
				rNodes = append(rNodes, defNode)
			}
			defNode.Repos = append(defNode.Repos, newrepos...)
		}
		ctx.ProgressBar().Suffix(fmt.Sprint(" ", 80, "%"))
		ctx.ProgressBar().Progress(80)

		// 5、将远程节点树合并入本地节点树。
		store.ReplaceNodes(rNodes)
		ctx.ProgressBar().Suffix(fmt.Sprint(" ", 100, "%"))
		ctx.ProgressBar().Progress(100)

		b, _ := view.GenJSONData(store.ListNodes())
		log.Debug().Msgf("Nodes: %s", b)

		ctx.ProgressBar().Stop()
	},
}
