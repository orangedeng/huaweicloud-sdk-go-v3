/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateWholeImageRequest struct {
	Body *CreateWholeImageRequestBody `json:"body,omitempty"`
}

func (o CreateWholeImageRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateWholeImageRequest", string(data)}, " ")
}
