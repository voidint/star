package gitee

import "github.com/voidint/star/plugin"

const (
	name = "gitee"
)

func init() {
	plugin.Register(name, new(holder))
}

type holder struct {
	token string
}

func (h *holder) Whoami() string {
	return name
}

func (h *holder) SetToken(token string) {
	h.token = token
}
