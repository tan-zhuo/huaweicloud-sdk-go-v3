package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopAutoCreateSnapshotsResponse Response Object
type StopAutoCreateSnapshotsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopAutoCreateSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopAutoCreateSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"StopAutoCreateSnapshotsResponse", string(data)}, " ")
}
