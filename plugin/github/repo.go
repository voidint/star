package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/voidint/star/plugin"
)

// GetRepo 返回认证用户下的仓库
// 参考 https://developer.github.com/v3/repos/#get
func (h *holder) GetRepo() (*plugin.Repo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", rootEndpoint, h.user.Login, h.conf.Repo)
	req, err := h.reqWithAuth(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("URL", url).Msg("Get repo")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, plugin.ErrRepoNotFound
	}

	var r plugin.Repo
	return &r, json.NewDecoder(resp.Body).Decode(&r)
}

// CreateRepo 为认证用户创建仓库
func (h *holder) CreateRepo() (*plugin.Repo, error) {
	url := fmt.Sprintf("%s/user/repos", rootEndpoint)

	reqBody := h.jsonBody4CreateRepo()
	req, err := h.reqWithAuth(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	log.Debug().Str("URL", url).Str("Request Body", string(reqBody)).Msg("Create repo")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnprocessableEntity {
		return nil, plugin.ErrCreateDuplicateRepo
	} else if resp.StatusCode == http.StatusNotFound {
		return nil, plugin.ErrPermissionDenied
	}
	var repo plugin.Repo
	return &repo, json.NewDecoder(resp.Body).Decode(&repo)
}

func (h *holder) jsonBody4CreateRepo() (data []byte) {
	data, _ = json.Marshal(map[string]interface{}{
		"name":        h.conf.Repo,
		"description": "A curated list of awesome frameworks, libraries, software and resources.",
		"private":     false,
		"auto_init":   true,
	})
	return data
}

// FetchAllStarredRepos 拉取当前用户标星的仓库列表。
func (h *holder) FetchAllStarredRepos() (items []*plugin.Repo, err error) {
	pg := &plugin.Pagination{
		Page:    1,
		PerPage: 100,
	}
	var repos []*plugin.Repo
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
func (h *holder) FetchStarredRepos(pg *plugin.Pagination) (items []*plugin.Repo, err error) {
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
