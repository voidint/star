package view

import (
	"fmt"
	"strings"

	"github.com/voidint/star/store"
)

// GenReadme 生成markdown格式的README内容。
// func GenReadme(tags []store.RepoedTag) (content []byte, err error) {
// 	var sb strings.Builder
// 	sb.WriteString("# Awesome\n\nThis file is autogenerated by [star](https://github.com/voidint/star), do not edit.\n\n")
// 	for i := range tags {
// 		sb.WriteString(fmt.Sprintf("## %s\n\n%s", tags[i].Tag.Name, tags[i].Tag.Desc))

// 		for j := range tags[i].Repos {
// 			sb.WriteString(fmt.Sprintf("- [%s](%s) %s\n",
// 				tags[i].Repos[j].FullName,
// 				tags[i].Repos[j].HTMLURL,
// 				tags[i].Repos[j].Description,
// 			))
// 		}
// 	}
// 	return []byte(sb.String()), nil
// }

// GenReadme 生成markdown格式的README内容。
func GenReadme(nodes []*store.Node) (content []byte, err error) {
	if len(nodes) <= 0 {
		return nil, nil
	}
	var sb strings.Builder
	sb.WriteString("# Awesome\n\nThis file is autogenerated by [star](https://github.com/voidint/star), do not edit.\n\n")
	for i := range nodes {
		sb.WriteString(fmt.Sprintf("## %s\n\n%s", nodes[i].Tag.Name, nodes[i].Tag.Desc))

		for j := range nodes[i].Repos {
			sb.WriteString(fmt.Sprintf("- [%s](%s) %s\n",
				nodes[i].Repos[j].FullName,
				nodes[i].Repos[j].HTMLURL,
				nodes[i].Repos[j].Description,
			))
		}

		b, err := GenReadme(nodes[i].Children)
		if err != nil {
			return nil, err
		}
		if b != nil {
			sb.WriteByte('\n')
			sb.Write(b)
		}
	}

	return []byte(sb.String()), nil
}
