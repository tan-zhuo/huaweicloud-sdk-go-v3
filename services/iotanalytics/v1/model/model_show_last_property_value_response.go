package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLastPropertyValueResponse Response Object
type ShowLastPropertyValueResponse struct {

	// 查询到的资产属性列表
	Properties     *[]AssetPropertyLastValue `json:"properties,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowLastPropertyValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLastPropertyValueResponse struct{}"
	}

	return strings.Join([]string{"ShowLastPropertyValueResponse", string(data)}, " ")
}
