package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/voidint/star/plugin"
)

// FetchAllStarredRepos 拉取当前用户标星的仓库列表。
func (h *holder) FetchAllStarredRepos() (items []*plugin.StarredRepo, err error) {
	pg := &plugin.Pagination{
		Page:    1,
		PerPage: 100,
	}
	var repos []*plugin.StarredRepo
	repos, err = h.FetchStarredRepos(pg)
	if err != nil {
		return nil, err
	}

	for len(repos) > 0 {
		items = append(items, repos...)

		pg.Page++
		repos, err = h.FetchStarredRepos(pg)
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}

// FetchStarredRepos 拉取当前用户标星的仓库分页列表。
// 参考 https://developer.github.com/v3/activity/starring/#list-repositories-being-starred
func (h *holder) FetchStarredRepos(pg *plugin.Pagination) (items []*plugin.StarredRepo, err error) {
	page, size := 1, 30
	if pg != nil && pg.Page > 0 {
		page = pg.Page
	}
	if pg != nil && pg.PerPage > 0 && pg.PerPage <= 100 {
		size = pg.PerPage
	}
	url := fmt.Sprintf("%s/user/starred?page=%d&per_page=%d", rootEndpoint, page, size)
	req, err := h.reqWithAuth(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	log.Debug().Str("URL", url).Msg("Fetch repos being starred")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return items, json.NewDecoder(resp.Body).Decode(&items)
}
