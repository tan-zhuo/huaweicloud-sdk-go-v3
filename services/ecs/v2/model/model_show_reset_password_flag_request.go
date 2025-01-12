package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResetPasswordFlagRequest Request Object
type ShowResetPasswordFlagRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowResetPasswordFlagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResetPasswordFlagRequest struct{}"
	}

	return strings.Join([]string{"ShowResetPasswordFlagRequest", string(data)}, " ")
}
