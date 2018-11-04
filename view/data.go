package view

import (
	"encoding/json"

	"github.com/voidint/star/store"
)

// GenJSONData 生成json格式的data.json内容。
// func GenJSONData(tags []store.RepoedTag) (content []byte, err error) {
// 	return json.MarshalIndent(tags, "  ", "  ")
// }

// GenJSONData 生成json格式的data.json内容。
func GenJSONData(nodes []*store.Node) (content []byte, err error) {
	return json.MarshalIndent(nodes, "  ", "  ")
}
