package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableControlRequest Request Object
type DisableControlRequest struct {
	Body *ControlOperateReqBody `json:"body,omitempty"`
}

func (o DisableControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableControlRequest struct{}"
	}

	return strings.Join([]string{"DisableControlRequest", string(data)}, " ")
}
