package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropertyValue struct {
}

func (o PropertyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyValue struct{}"
	}

	return strings.Join([]string{"PropertyValue", string(data)}, " ")
}
