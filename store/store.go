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

	stores = make(map[string][]*Node)
	for _, hName := range plugin.Registered() {
		log.Debug().Msgf("init store for %s", hName)
		stores[hName] = []*Node{ // 初始化一个带默认tag的节点
			{
				Tag:      &tag.DefaultTag,
				Repos:    []*plugin.Repo{},
				Children: []*Node{},
			},
		}
	}
}

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

// Node 节点
type Node struct {
	Tag      *tag.Tag       `json:"tag"`
	Repos    []*plugin.Repo `json:"repos"`
	Children []*Node        `json:"children"`
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

// ReplaceNodes 替换当前常见的节点列表
func ReplaceNodes(nodes []*Node) {
	once.Do(func() {
		initStores()
	})

	hmux.Lock()
	defer hmux.Unlock()

	stores[curHolder] = nodes
}
