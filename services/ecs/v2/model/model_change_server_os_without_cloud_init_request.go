package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeServerOsWithoutCloudInitRequest struct {
	// 云服务器ID。

	ServerId string `json:"server_id"`

	Body *ChangeServerOsWithoutCloudInitRequestBody `json:"body,omitempty"`
}

func (o ChangeServerOsWithoutCloudInitRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithoutCloudInitRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithoutCloudInitRequest", string(data)}, " ")
}
