package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemuxTaskResponse Response Object
type ListRemuxTaskResponse struct {

	// 任务总数
	Total *int32 `json:"total,omitempty"`

	// 任务列表
	Tasks          *[]RemuxTask `json:"tasks,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRemuxTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"ListRemuxTaskResponse", string(data)}, " ")
}
