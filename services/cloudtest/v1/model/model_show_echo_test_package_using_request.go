package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEchoTestPackageUsingRequest Request Object
type ShowEchoTestPackageUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`
}

func (o ShowEchoTestPackageUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEchoTestPackageUsingRequest struct{}"
	}

	return strings.Join([]string{"ShowEchoTestPackageUsingRequest", string(data)}, " ")
}
