package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateipRequest Request Object
type ShowPrivateipRequest struct {

	// 私有IP ID
	PrivateipId string `json:"privateip_id"`
}

func (o ShowPrivateipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateipRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateipRequest", string(data)}, " ")
}
