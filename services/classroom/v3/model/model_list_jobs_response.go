package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsResponse Response Object
type ListJobsResponse struct {

	// 作业列表
	Jobs *[]JobCard `json:"jobs,omitempty"`

	// 作业总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsResponse struct{}"
	}

	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}
