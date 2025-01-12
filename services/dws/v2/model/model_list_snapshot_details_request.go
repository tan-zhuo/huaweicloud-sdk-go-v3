package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotDetailsRequest Request Object
type ListSnapshotDetailsRequest struct {

	// 快照ID。
	SnapshotId string `json:"snapshot_id"`
}

func (o ListSnapshotDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotDetailsRequest", string(data)}, " ")
}
