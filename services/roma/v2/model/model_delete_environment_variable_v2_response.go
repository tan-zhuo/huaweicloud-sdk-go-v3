package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnvironmentVariableV2Response Response Object
type DeleteEnvironmentVariableV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEnvironmentVariableV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentVariableV2Response struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentVariableV2Response", string(data)}, " ")
}
