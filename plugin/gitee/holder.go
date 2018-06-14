package gitee

const (
	name = "gitee"
)

// func init() {
// 	plugin.Register(name, new(holder))
// }

type holder struct {
	token string
}

func (h *holder) Whoami() string {
	return name
}
