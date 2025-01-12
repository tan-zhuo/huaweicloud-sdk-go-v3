package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAnimatedGraphicsTaskResponse Response Object
type ListAnimatedGraphicsTaskResponse struct {

	// 任务总数
	Total *int32 `json:"total,omitempty"`

	// 任务列表
	Tasks          *[]AnimatedGraphicsTask `json:"tasks,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListAnimatedGraphicsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"ListAnimatedGraphicsTaskResponse", string(data)}, " ")
}
