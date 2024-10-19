package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotificationRequest Request Object
type CreateNotificationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateNotificationRequestBody `json:"body,omitempty"`
}

func (o CreateNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationRequest struct{}"
	}

	return strings.Join([]string{"CreateNotificationRequest", string(data)}, " ")
}
