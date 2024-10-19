package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogLtsConfigsResponse Response Object
type DeleteLogLtsConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLogLtsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogLtsConfigsResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogLtsConfigsResponse", string(data)}, " ")
}
