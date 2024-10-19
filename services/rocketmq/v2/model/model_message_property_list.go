package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MessagePropertyList struct {

	// 属性名称。
	Name *string `json:"name,omitempty"`

	// 属性值。
	Value *string `json:"value,omitempty"`
}

func (o MessagePropertyList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessagePropertyList struct{}"
	}

	return strings.Join([]string{"MessagePropertyList", string(data)}, " ")
}
