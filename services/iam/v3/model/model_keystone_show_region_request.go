package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowRegionRequest struct {
	RegionId string `json:"region_id"`
}

func (o KeystoneShowRegionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowRegionRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowRegionRequest", string(data)}, " ")
}
