package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Selector 选择器
type Selector struct {
	FieldSelector *FieldSelector `json:"fieldSelector,omitempty"`
}

func (o Selector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Selector struct{}"
	}

	return strings.Join([]string{"Selector", string(data)}, " ")
}
