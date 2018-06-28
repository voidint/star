package store

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/voidint/star/plugin"
	"github.com/voidint/star/tag"
)

var (
	once      sync.Once
	stores    map[string]*store
	hmux      sync.Mutex
	curStore  *store
	curHolder string
)

func initStores() {
	hmux.Lock()
	defer hmux.Unlock()
	log.Debug().Msg("init stores")

	stores = make(map[string]*store)
	for _, hName := range plugin.Registered() {
		log.Debug().Msgf("init store for %s", hName)
		stores[hName] = &store{
			tags: []*RepoedTag{
				{Tag: tag.Default()},
			},
		}
	}
}

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

// Use 切换到指定Holder
func Use(holderName string) {
	once.Do(func() {
		initStores()
	})

	hmux.Lock()
	defer hmux.Unlock()

	for k, v := range stores {
		if k == holderName {
			curHolder = k
			curStore = v
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
