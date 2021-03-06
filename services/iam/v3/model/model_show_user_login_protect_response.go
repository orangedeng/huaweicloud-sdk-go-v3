package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowUserLoginProtectResponse struct {
	LoginProtect   *LoginProtectResult `json:"login_protect,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowUserLoginProtectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserLoginProtectResponse struct{}"
	}

	return strings.Join([]string{"ShowUserLoginProtectResponse", string(data)}, " ")
}
