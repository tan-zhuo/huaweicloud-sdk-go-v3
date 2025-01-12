package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNicsRequest Request Object
type DeleteNicsRequest struct {

	// 边缘实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteNicsRequestBody `json:"body,omitempty"`
}

func (o DeleteNicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNicsRequest struct{}"
	}

	return strings.Join([]string{"DeleteNicsRequest", string(data)}, " ")
}
