package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCloudPhonePropertyRequest Request Object
type UpdateCloudPhonePropertyRequest struct {
	Body *UpdateCloudPhonePropertyRequestBody `json:"body,omitempty"`
}

func (o UpdateCloudPhonePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudPhonePropertyRequest struct{}"
	}

	return strings.Join([]string{"UpdateCloudPhonePropertyRequest", string(data)}, " ")
}
