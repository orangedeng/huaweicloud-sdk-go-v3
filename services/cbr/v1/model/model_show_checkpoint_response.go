/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCheckpointResponse struct {
	Checkpoint *CheckpointCreate `json:"checkpoint,omitempty"`
}

func (o ShowCheckpointResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCheckpointResponse", string(data)}, " ")
}
