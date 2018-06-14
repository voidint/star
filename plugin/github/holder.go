package github

import "github.com/voidint/star/plugin"

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
	defer func() {
		if u != nil && err == nil {
			h.auth = auth
			h.user = u
		}
	}()
	if auth.Token != "" {
		return h.fetchUserByToken(auth.Token)
	}
	return h.fetchUserByBasic(auth.Username, auth.Password)
}

func (h *holder) Init() error {
	// stars, err := h.fetchAllStars()
	// if err != nil {
	// 	return err
	// }
	return nil
}
