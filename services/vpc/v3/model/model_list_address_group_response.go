package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressGroupResponse Response Object
type ListAddressGroupResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 地址组列表响应体
	AddressGroups *[]AddressGroup `json:"address_groups,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"ListAddressGroupResponse", string(data)}, " ")
}
