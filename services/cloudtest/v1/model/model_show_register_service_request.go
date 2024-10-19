package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRegisterServiceRequest Request Object
type ShowRegisterServiceRequest struct {
}

func (o ShowRegisterServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRegisterServiceRequest struct{}"
	}

	return strings.Join([]string{"ShowRegisterServiceRequest", string(data)}, " ")
}
