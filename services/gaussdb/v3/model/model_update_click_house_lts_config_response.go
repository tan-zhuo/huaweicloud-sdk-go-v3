package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClickHouseLtsConfigResponse Response Object
type UpdateClickHouseLtsConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateClickHouseLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClickHouseLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateClickHouseLtsConfigResponse", string(data)}, " ")
}
