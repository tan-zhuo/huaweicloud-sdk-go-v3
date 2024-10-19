package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmConfigRequest Request Object
type ShowAlarmConfigRequest struct {
}

func (o ShowAlarmConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmConfigRequest", string(data)}, " ")
}
