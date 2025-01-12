package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProcessRequest Request Object
type ShowProcessRequest struct {
}

func (o ShowProcessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProcessRequest struct{}"
	}

	return strings.Join([]string{"ShowProcessRequest", string(data)}, " ")
}
