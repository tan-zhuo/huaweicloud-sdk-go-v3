package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackendInstanceV2Response Response Object
type DeleteBackendInstanceV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackendInstanceV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackendInstanceV2Response struct{}"
	}

	return strings.Join([]string{"DeleteBackendInstanceV2Response", string(data)}, " ")
}
