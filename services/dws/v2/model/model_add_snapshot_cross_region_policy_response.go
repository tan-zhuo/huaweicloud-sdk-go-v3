package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSnapshotCrossRegionPolicyResponse Response Object
type AddSnapshotCrossRegionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddSnapshotCrossRegionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSnapshotCrossRegionPolicyResponse struct{}"
	}

	return strings.Join([]string{"AddSnapshotCrossRegionPolicyResponse", string(data)}, " ")
}
