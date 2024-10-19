package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAddressSetResponse Response Object
type AddAddressSetResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddAddressSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressSetResponse struct{}"
	}

	return strings.Join([]string{"AddAddressSetResponse", string(data)}, " ")
}
