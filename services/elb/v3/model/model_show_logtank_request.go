package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogtankRequest Request Object
type ShowLogtankRequest struct {

	// 云日志ID。
	LogtankId string `json:"logtank_id"`
}

func (o ShowLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogtankRequest struct{}"
	}

	return strings.Join([]string{"ShowLogtankRequest", string(data)}, " ")
}
