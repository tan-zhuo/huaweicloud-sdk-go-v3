package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAddressGroupResponse Response Object
type CreateAddressGroupResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	AddressGroup   *AddressGroup `json:"address_group,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateAddressGroupResponse", string(data)}, " ")
}
