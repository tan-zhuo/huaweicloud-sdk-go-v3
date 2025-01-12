package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBindingDeviceRequest Request Object
type CreateBindingDeviceRequest struct {
	Body *BindMfaDevice `json:"body,omitempty"`
}

func (o CreateBindingDeviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBindingDeviceRequest struct{}"
	}

	return strings.Join([]string{"CreateBindingDeviceRequest", string(data)}, " ")
}
