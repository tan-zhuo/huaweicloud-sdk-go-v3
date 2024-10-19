package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvRequest Request Object
type ShowEnvRequest struct {
}

func (o ShowEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvRequest struct{}"
	}

	return strings.Join([]string{"ShowEnvRequest", string(data)}, " ")
}
