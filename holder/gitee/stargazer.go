package gitee

import (
	"github.com/voidint/star/holder"
)

const (
	name = "gitee"
)

func init() {
	holder.Register(name, new(stargazer))
}

type stargazer struct {
}

func (ctrl *stargazer) Whoami() string {
	return name
}
