package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReplicationJobsResponse Response Object
type ListReplicationJobsResponse struct {

	// 备份迁移任务总数。
	Count *int32 `json:"count,omitempty"`

	// 备份迁移任务列表。
	Jobs           *[]OfflineTaskInfo `json:"jobs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListReplicationJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationJobsResponse struct{}"
	}

	return strings.Join([]string{"ListReplicationJobsResponse", string(data)}, " ")
}
