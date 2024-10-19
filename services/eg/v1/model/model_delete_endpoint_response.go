package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointResponse Response Object
type DeleteEndpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointResponse", string(data)}, " ")
}
