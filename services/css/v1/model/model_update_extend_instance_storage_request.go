package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExtendInstanceStorageRequest Request Object
type UpdateExtendInstanceStorageRequest struct {

	// 指定待扩容的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *RoleExtendReq `json:"body,omitempty"`
}

func (o UpdateExtendInstanceStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExtendInstanceStorageRequest struct{}"
	}

	return strings.Join([]string{"UpdateExtendInstanceStorageRequest", string(data)}, " ")
}
