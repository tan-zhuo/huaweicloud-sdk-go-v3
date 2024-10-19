package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterVolumesRequest Request Object
type ListLogicalClusterVolumesRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 分页查询，偏移
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询，每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListLogicalClusterVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterVolumesRequest struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterVolumesRequest", string(data)}, " ")
}
