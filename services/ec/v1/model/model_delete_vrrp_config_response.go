package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVrrpConfigResponse Response Object
type DeleteVrrpConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVrrpConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVrrpConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteVrrpConfigResponse", string(data)}, " ")
}
