package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowEndpointRequest struct {
	EndpointId string `json:"endpoint_id"`
}

func (o KeystoneShowEndpointRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowEndpointRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowEndpointRequest", string(data)}, " ")
}
