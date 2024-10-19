package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneProperty 云手机属性信息。
type PhoneProperty struct {

	// 云手机id。
	PhoneId string `json:"phone_id"`

	// 云手机属性列表，为Json格式字符串。
	Property *string `json:"property,omitempty"`
}

func (o PhoneProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneProperty struct{}"
	}

	return strings.Join([]string{"PhoneProperty", string(data)}, " ")
}
