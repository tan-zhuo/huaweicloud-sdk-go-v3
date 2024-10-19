package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVrrpConfigRequest Request Object
type ShowVrrpConfigRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`
}

func (o ShowVrrpConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVrrpConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowVrrpConfigRequest", string(data)}, " ")
}
