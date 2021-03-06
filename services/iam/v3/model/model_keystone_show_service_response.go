package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowServiceResponse struct {
	Service        *Service `json:"service,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o KeystoneShowServiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowServiceResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowServiceResponse", string(data)}, " ")
}
