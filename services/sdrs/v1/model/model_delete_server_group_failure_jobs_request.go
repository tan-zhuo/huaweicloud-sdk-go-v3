package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerGroupFailureJobsRequest Request Object
type DeleteServerGroupFailureJobsRequest struct {

	// 保护组ID。
	ServerGroupId string `json:"server_group_id"`
}

func (o DeleteServerGroupFailureJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupFailureJobsRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupFailureJobsRequest", string(data)}, " ")
}
