package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePredefineTagsRequest struct {
	Body *ModifyPrefineTag `json:"body,omitempty"`
}

func (o UpdatePredefineTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePredefineTagsRequest struct{}"
	}

	return strings.Join([]string{"UpdatePredefineTagsRequest", string(data)}, " ")
}
