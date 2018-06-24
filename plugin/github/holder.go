package github

import (
	"fmt"
	"io"
	"net/http"

	"github.com/voidint/star/plugin"
)

const (
	name         = "github"
	rootEndpoint = "https://api.github.com"
)

func init() {
	plugin.Register(name, &holder{
		conf: &plugin.Configuration{
			Repo:   plugin.DefRepo,
			Branch: plugin.DefBranch,
		},
	})
}

type holder struct {
	auth *plugin.Authentication
	user *plugin.User
	conf *plugin.Configuration
}

func (h *holder) Whoami() string {
	return name
}

// SetConfiguration 设置配置参数
func (h *holder) SetConfiguration(conf *plugin.Configuration) {
	h.conf.Repo = conf.Repo
	h.conf.Branch = conf.Branch
}

// Login 登录认证并返回用户信息。若认证信息有误，则返回plugin.ErrBadCredentials。
func (h *holder) Login(auth *plugin.Authentication) (u *plugin.User, err error) {
	if auth == nil {
		return nil, plugin.ErrBadCredentials
	}
	h.auth = auth

	defer func() {
		if u != nil && err == nil {
			h.user = u
		} else {
			h.auth = nil
		}
	}()

	return h.fetchUser()
}

func (h *holder) reqWithAuth(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if h.auth == nil {
		return nil, plugin.ErrBadCredentials
	}
	if h.auth.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("token %s", h.auth.Token))
	} else {
		req.SetBasicAuth(h.auth.Username, h.auth.Password)
	}
	return req, nil
}
