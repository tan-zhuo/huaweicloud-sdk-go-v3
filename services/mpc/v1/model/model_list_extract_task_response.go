package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExtractTaskResponse Response Object
type ListExtractTaskResponse struct {

	// 任务总数
	Total *int32 `json:"total,omitempty"`

	// 任务列表
	Tasks          *[]ExtractTask `json:"tasks,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListExtractTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtractTaskResponse struct{}"
	}

	return strings.Join([]string{"ListExtractTaskResponse", string(data)}, " ")
}
