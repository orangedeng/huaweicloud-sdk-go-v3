package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBlockchainsRequest struct {
}

func (o ListBlockchainsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBlockchainsRequest struct{}"
	}

	return strings.Join([]string{"ListBlockchainsRequest", string(data)}, " ")
}
