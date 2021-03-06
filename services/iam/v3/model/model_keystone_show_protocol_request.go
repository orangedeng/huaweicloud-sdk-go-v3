package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowProtocolRequest struct {
	IdpId      string `json:"idp_id"`
	ProtocolId string `json:"protocol_id"`
}

func (o KeystoneShowProtocolRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowProtocolRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowProtocolRequest", string(data)}, " ")
}
