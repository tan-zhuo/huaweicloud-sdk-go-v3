package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyItemCondition struct {

	// 条件类型
	Type *string `json:"type,omitempty"`

	// 条件值
	Values *[]string `json:"values,omitempty"`
}

func (o PolicyItemCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyItemCondition struct{}"
	}

	return strings.Join([]string{"PolicyItemCondition", string(data)}, " ")
}
