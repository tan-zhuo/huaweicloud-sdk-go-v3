package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddNicsRequest Request Object
type AddNicsRequest struct {

	// 边缘实例ID。
	InstanceId string `json:"instance_id"`

	Body *AddNicsRequestBody `json:"body,omitempty"`
}

func (o AddNicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNicsRequest struct{}"
	}

	return strings.Join([]string{"AddNicsRequest", string(data)}, " ")
}
