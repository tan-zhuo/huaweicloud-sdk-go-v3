package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAddressGroupResponse Response Object
type UpdateAddressGroupResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	AddressGroup   *AddressGroup `json:"address_group,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateAddressGroupResponse", string(data)}, " ")
}
