package store

import (
	"errors"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/voidint/star/plugin"
	"github.com/voidint/star/tag"
)

var (
	// stores    map[string]*store  // 每个插件一个store
	// curStore  *store
	once      sync.Once
	hmux      sync.Mutex
	stores    map[string][]*Node
	curHolder string
)

func initStores() {
	hmux.Lock()
	defer hmux.Unlock()
	log.Debug().Msg("init stores")

	// stores = make(map[string]*store)
	// for _, hName := range plugin.Registered() {
	// 	log.Debug().Msgf("init store for %s", hName)
	// 	stores[hName] = &store{
	// 		tags: []*RepoedTag{
	// 			{Tag: tag.Default()},
	// 		},
	// 	}
	// }
	stores = make(map[string][]*Node)
	for _, hName := range plugin.Registered() {
		log.Debug().Msgf("init store for %s", hName)
		stores[hName] = []*Node{ // 初始化一个带默认tag的节点
			{
				Tag:      tag.Default(),
				Repos:    []*plugin.Repo{},
				Children: []*Node{},
			},
		}
	}
}

/*
type store struct {
	tags  []*RepoedTag
	repos []*TaggedRepo
}

// RepoedTag 包含标星仓库列表的Tag
type RepoedTag struct {
	Tag   *tag.Tag
	Repos []*plugin.Repo
}

// TaggedRepo 包含Tag列表的标星仓库
type TaggedRepo struct {
	Repo *plugin.Repo
	Tags []*tag.Tag
}
*/

// Use 切换到指定Holder
func Use(holderName string) {
	once.Do(func() {
		initStores()
	})

	hmux.Lock()
	defer hmux.Unlock()

	for k := range stores {
		if k == holderName {
			curHolder = k
			// curStore = v
			return
		}
	}
	panic("unregistered holder name")
}

// CurrentHolder 返回当前Holder名称
func CurrentHolder() string {
	once.Do(func() {
		initStores()
	})

	hmux.Lock()
	defer hmux.Unlock()
	return curHolder
}

/*
// OverwriteRepoedTag 覆写当前Holder的标星仓库
func OverwriteRepoedTag(tagName string, repos []*plugin.Repo) {
	once.Do(func() {
		initStores()
	})

	hmux.Lock()
	defer hmux.Unlock()

	curStore.tags = []*RepoedTag{
		{
			Tag:   tag.Default(), // TODO find tag by name
			Repos: repos,
		},
	}
}

// ListRepoedTag 返回当前Holder的标星仓库列表
func ListRepoedTag() []RepoedTag {
	once.Do(func() {
		initStores()
	})

	hmux.Lock()
	defer hmux.Unlock()

	items := make([]RepoedTag, 0, len(curStore.tags))
	for i := range curStore.tags {
		items = append(items, *curStore.tags[i])
	}
	return items
}

// TagRepo 为标星仓库打标签
func TagRepo(tag, repo string) error {
	return nil
}
*/

// Node 节点
type Node struct {
	Tag      *tag.Tag
	Repos    []*plugin.Repo
	Children []*Node
}

// Nodes 节点集合
type Nodes []*Node

// FindNode 查找指定tag路径的节点。典型的tag路径如'/计算机/网络/tcp'
func (nodes Nodes) FindNode(tagPath string) *Node {
	for _, node := range nodes {
		if node.Tag.Path == tagPath {
			return node
		}
		if fnode := Nodes(node.Children).FindNode(tagPath); fnode != nil {
			return fnode
		}
	}
	return nil
}

var (
	// ErrNodeNotFound 节点不存在
	ErrNodeNotFound = errors.New("node not found")
)

// OverwriteRepos 覆盖指定路径标签的仓库列表
func OverwriteRepos(tagPath string, repos []*plugin.Repo) error {
	once.Do(func() {
		initStores()
	})

	hmux.Lock()
	defer hmux.Unlock()

	nodes := stores[curHolder]

	fnode := Nodes(nodes).FindNode(tagPath)
	if fnode == nil {
		return ErrNodeNotFound
	}
	fnode.Repos = repos
	return nil
}

// ListNodes 节点列表
func ListNodes() []*Node {
	once.Do(func() {
		initStores()
	})

	hmux.Lock()
	defer hmux.Unlock()

	return stores[curHolder]
}

// FindNode 查找指定tag路径的节点。典型的tag路径如'/计算机/网络/tcp'
// func (nodes Nodes) FindNode(tagPath string) *Node {
// 	tags := strings.Split(tagPath, "/")
// 	// TODO 清除空的tag元素

// 	for i, tag := range tags {
// 		if tag == "" {
// 			continue
// 		}
// 		for _, node := range nodes {
// 			if node.Tag != nil && node.Tag.Name == tagPath {
// 				if i == len(tags)-1 {
// 					return node
// 				}
// 				findNodeByTagName(node.Children)
// 			}

// 		}
// 	}
// }

// func findNodeByTagName(nodes []*Node, tagName string) *Node {
// 	for i := range nodes {
// 		if nodes[i].Tag != nil && nodes[i].Tag.Name == tagName {
// 			return nodes[i]
// 		}
// 	}
// 	return nil
// }
