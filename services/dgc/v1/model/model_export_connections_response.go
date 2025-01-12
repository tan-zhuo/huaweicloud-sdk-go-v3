package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportConnectionsResponse Response Object
type ExportConnectionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ExportConnectionsResponse", string(data)}, " ")
}
