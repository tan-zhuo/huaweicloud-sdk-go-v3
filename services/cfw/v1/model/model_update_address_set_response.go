package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAddressSetResponse Response Object
type UpdateAddressSetResponse struct {
	Data           *UpdateAddressSetResponseData `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o UpdateAddressSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressSetResponse struct{}"
	}

	return strings.Join([]string{"UpdateAddressSetResponse", string(data)}, " ")
}
