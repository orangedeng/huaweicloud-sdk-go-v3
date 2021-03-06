package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateLoginProjectReq struct {
	LoginProtect *UpdateLoginProject `json:"login_protect"`
}

func (o UpdateLoginProjectReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoginProjectReq struct{}"
	}

	return strings.Join([]string{"UpdateLoginProjectReq", string(data)}, " ")
}
