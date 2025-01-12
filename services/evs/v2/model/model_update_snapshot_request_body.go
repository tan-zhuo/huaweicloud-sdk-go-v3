package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSnapshotRequestBody This is a auto create Body Object
type UpdateSnapshotRequestBody struct {
	Snapshot *UpdateSnapshotOption `json:"snapshot"`
}

func (o UpdateSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotRequestBody", string(data)}, " ")
}
