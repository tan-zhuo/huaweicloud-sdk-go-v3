package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvironmentDetailRequest Request Object
type ShowEnvironmentDetailRequest struct {

	// 环境ID。
	EnvironmentId string `json:"environment_id"`
}

func (o ShowEnvironmentDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvironmentDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowEnvironmentDetailRequest", string(data)}, " ")
}
