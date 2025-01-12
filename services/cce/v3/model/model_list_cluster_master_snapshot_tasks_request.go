package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterMasterSnapshotTasksRequest Request Object
type ListClusterMasterSnapshotTasksRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterMasterSnapshotTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterMasterSnapshotTasksRequest struct{}"
	}

	return strings.Join([]string{"ListClusterMasterSnapshotTasksRequest", string(data)}, " ")
}
