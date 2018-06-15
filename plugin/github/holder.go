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
	plugin.Register(name, new(holder))
}

type holder struct {
	auth *plugin.Authentication
	user *plugin.User
}

func (h *holder) Whoami() string {
	return name
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

func (h *holder) Init() error {
	// stars, err := h.fetchAllStars()
	// if err != nil {
	// 	return err
	// }
	return nil
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
