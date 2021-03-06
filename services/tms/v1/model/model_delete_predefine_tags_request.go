package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePredefineTagsRequest struct {
	Body *ReqDeletePredefineTag `json:"body,omitempty"`
}

func (o DeletePredefineTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePredefineTagsRequest struct{}"
	}

	return strings.Join([]string{"DeletePredefineTagsRequest", string(data)}, " ")
}
