package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterRequest Request Object
type ResizeClusterRequest struct {

	// 待调整大小的集群ID
	ClusterId string `json:"cluster_id"`

	Body *ResizeClusterRequestBody `json:"body,omitempty"`
}

func (o ResizeClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterRequest struct{}"
	}

	return strings.Join([]string{"ResizeClusterRequest", string(data)}, " ")
}
