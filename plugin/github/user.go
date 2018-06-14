package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/voidint/star/plugin"
)

// 参考 https://developer.github.com/v3/users/#get-the-authenticated-user
func (h *holder) fetchUserByToken(token string) (*plugin.User, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user", rootEndpoint), nil)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, plugin.ErrBadCredentials
	}

	var user plugin.User
	return &user, json.NewDecoder(resp.Body).Decode(&user)
}

// 参考 https://developer.github.com/v3/users/#get-the-authenticated-user
func (h *holder) fetchUserByBasic(username, password string) (*plugin.User, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user", rootEndpoint), nil)
	req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, plugin.ErrBadCredentials
	}

	var user plugin.User
	return &user, json.NewDecoder(resp.Body).Decode(&user)
}
