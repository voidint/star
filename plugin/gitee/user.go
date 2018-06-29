package gitee

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/voidint/star/plugin"
)

// 参考 https://gitee.com/api/v5/swagger#/getV5User
func (h *holder) fetchUser() (*plugin.User, error) {
	req, _ := h.reqWithAuth(http.MethodGet, fmt.Sprintf("%s/user", rootEndpoint), nil)

	log.Debug().Str("GET URL", req.URL.String()).Msg("Get user")
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
