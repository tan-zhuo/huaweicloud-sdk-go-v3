package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAddressItemResponse Response Object
type DeleteAddressItemResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteAddressItemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressItemResponse struct{}"
	}

	return strings.Join([]string{"DeleteAddressItemResponse", string(data)}, " ")
}
