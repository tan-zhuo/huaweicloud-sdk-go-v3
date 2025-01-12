package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceTemplateByIdResponse Response Object
type UpdateDeviceTemplateByIdResponse struct {
	DeviceTemplate *EdgemgrDevice `json:"device_template,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateDeviceTemplateByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceTemplateByIdResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceTemplateByIdResponse", string(data)}, " ")
}
