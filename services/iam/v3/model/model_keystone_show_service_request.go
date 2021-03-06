package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowServiceRequest struct {
	ServiceId string `json:"service_id"`
}

func (o KeystoneShowServiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowServiceRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowServiceRequest", string(data)}, " ")
}
