package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertConfigRequest Request Object
type ShowAlertConfigRequest struct {
}

func (o ShowAlertConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAlertConfigRequest", string(data)}, " ")
}
