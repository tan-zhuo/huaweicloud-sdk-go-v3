package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAllServerGroupFailureJobsResponse Response Object
type DeleteAllServerGroupFailureJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAllServerGroupFailureJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAllServerGroupFailureJobsResponse struct{}"
	}

	return strings.Join([]string{"DeleteAllServerGroupFailureJobsResponse", string(data)}, " ")
}
