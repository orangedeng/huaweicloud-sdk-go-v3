/*
 * cts
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
type UpdateTrackerRequest struct {
	Body *UpdateTrackerRequestBody `json:"body,omitempty"`
}

func (o UpdateTrackerRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateTrackerRequest", string(data)}, " ")
}
