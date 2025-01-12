package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotsResponse Response Object
type ListSnapshotsResponse struct {

	// 快照列表。
	Backups        *[]ListSnapshotBackupsResp `json:"backups,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotsResponse", string(data)}, " ")
}
