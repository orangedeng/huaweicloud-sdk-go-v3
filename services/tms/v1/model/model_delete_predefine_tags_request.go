/*
 * TMS
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
type DeletePredefineTagsRequest struct {
	Body *ReqDeletePredefineTag `json:"body,omitempty"`
}

func (o DeletePredefineTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePredefineTagsRequest", string(data)}, " ")
}
