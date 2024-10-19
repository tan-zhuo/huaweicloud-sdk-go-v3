package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAction struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 类型，默认AOP_WORKFLOW.
	ActionType string `json:"action_type"`

	// 剧本动作ID
	ActionId string `json:"action_id"`

	// 排序方式
	SortOrder *string `json:"sort_order,omitempty"`
}

func (o CreateAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAction struct{}"
	}

	return strings.Join([]string{"CreateAction", string(data)}, " ")
}
