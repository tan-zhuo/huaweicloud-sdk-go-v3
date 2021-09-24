package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {
	Server         *Instance `json:"server,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}