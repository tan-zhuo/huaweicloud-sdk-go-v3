package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotResponse Response Object
type CreateSnapshotResponse struct {
	Snapshot       *SnapshotResp `json:"snapshot,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotResponse struct{}"
	}

	return strings.Join([]string{"CreateSnapshotResponse", string(data)}, " ")
}
