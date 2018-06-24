package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/voidint/star/plugin"
)

// GetFile 获取文件内容。
// 参考 https://developer.github.com/v3/repos/contents/#get-contents
func (h *holder) GetFile(path string) (*plugin.File, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/contents/%s", rootEndpoint, h.auth.Username, h.conf.Repo, path)
	req, err := h.reqWithAuth(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("URL", url).Msg("Get file content")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, plugin.ErrFileNotFound
	}

	var f plugin.File
	return &f, json.NewDecoder(resp.Body).Decode(&f)
}

// CreateFile 在目标路径下创建文件。
// 参考 https://developer.github.com/v3/repos/contents/#create-a-file
func (h *holder) CreateFile(path string, content []byte) (*plugin.File, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/contents/%s", rootEndpoint, h.auth.Username, h.conf.Repo, path)

	reqBody := h.jsonBody4CreateFile(content)
	req, err := h.reqWithAuth(http.MethodPut, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	log.Debug().Str("URL", url).Str("Request Body", string(reqBody)).Msg("Create a file")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, plugin.ErrCreateFile
	}

	var respBody struct {
		Content plugin.File `json:"content"`
	}
	return &respBody.Content, json.NewDecoder(resp.Body).Decode(&respBody)
}

func (h *holder) jsonBody4CreateFile(content []byte) (data []byte) {
	data, _ = json.Marshal(map[string]interface{}{
		"message": "update a file",
		"content": base64.StdEncoding.EncodeToString(content),
		"branch":  h.conf.Branch,
	})
	return data
}

// UpdateFile 更新文件。
// 参考 https://developer.github.com/v3/repos/contents/#update-a-file
func (h *holder) UpdateFile(path, sha string, content []byte) (*plugin.File, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/contents/%s", rootEndpoint, h.auth.Username, h.conf.Repo, path)

	reqBody := h.jsonBody4UpdateFile(sha, content)
	req, err := h.reqWithAuth(http.MethodPut, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	log.Debug().Str("URL", url).Str("Request Body", string(reqBody)).Msg("Update a file")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, plugin.ErrUpdateFile
	}

	var respBody struct {
		Content plugin.File `json:"content"`
	}
	return &respBody.Content, json.NewDecoder(resp.Body).Decode(&respBody)
}

func (h *holder) jsonBody4UpdateFile(sha string, content []byte) (data []byte) {
	data, _ = json.Marshal(map[string]interface{}{
		"message": "update a file",
		"content": base64.StdEncoding.EncodeToString(content),
		"sha":     sha,
		"branch":  h.conf.Branch,
	})
	return data
}
