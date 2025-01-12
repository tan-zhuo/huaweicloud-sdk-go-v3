package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateipRequest Request Object
type DeletePrivateipRequest struct {

	// 私有IP ID
	PrivateipId string `json:"privateip_id"`
}

func (o DeletePrivateipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateipRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateipRequest", string(data)}, " ")
}
