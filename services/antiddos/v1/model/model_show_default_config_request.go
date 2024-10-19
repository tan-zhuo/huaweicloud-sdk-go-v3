package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultConfigRequest Request Object
type ShowDefaultConfigRequest struct {
}

func (o ShowDefaultConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowDefaultConfigRequest", string(data)}, " ")
}
