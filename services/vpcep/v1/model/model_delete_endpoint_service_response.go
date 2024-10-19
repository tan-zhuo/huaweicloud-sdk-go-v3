package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointServiceResponse Response Object
type DeleteEndpointServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"DeleteEndpointServiceResponse", string(data)}, " ")
}
