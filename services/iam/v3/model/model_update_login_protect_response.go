package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateLoginProtectResponse struct {
	LoginProtect   *UpdateLoginProtectRespon `json:"login_protect,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o UpdateLoginProtectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoginProtectResponse struct{}"
	}

	return strings.Join([]string{"UpdateLoginProtectResponse", string(data)}, " ")
}
