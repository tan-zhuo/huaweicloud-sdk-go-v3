package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueueDetailsRequest Request Object
type ShowQueueDetailsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 所属Vhost名称
	Vhost string `json:"vhost"`

	// Queue名称
	Queue string `json:"queue"`
}

func (o ShowQueueDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueDetailsRequest", string(data)}, " ")
}
