/*
 * IAM
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
type KeystoneShowVersionRequest struct {
}

func (o KeystoneShowVersionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneShowVersionRequest", string(data)}, " ")
}
