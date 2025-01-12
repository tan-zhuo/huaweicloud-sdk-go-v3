package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExtendInstanceStorageResponse Response Object
type UpdateExtendInstanceStorageResponse struct {

	// 集群ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateExtendInstanceStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExtendInstanceStorageResponse struct{}"
	}

	return strings.Join([]string{"UpdateExtendInstanceStorageResponse", string(data)}, " ")
}
