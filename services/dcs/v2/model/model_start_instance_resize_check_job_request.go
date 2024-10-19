package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInstanceResizeCheckJobRequest Request Object
type StartInstanceResizeCheckJobRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *StartInstanceResizeCheckJobRequestBody `json:"body,omitempty"`
}

func (o StartInstanceResizeCheckJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceResizeCheckJobRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceResizeCheckJobRequest", string(data)}, " ")
}
