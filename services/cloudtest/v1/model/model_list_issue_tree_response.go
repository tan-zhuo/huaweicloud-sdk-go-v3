package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueTreeResponse Response Object
type ListIssueTreeResponse struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value          *[]WorkItemVo `json:"value,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListIssueTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueTreeResponse struct{}"
	}

	return strings.Join([]string{"ListIssueTreeResponse", string(data)}, " ")
}
