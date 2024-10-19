package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachPluginToApiRequest Request Object
type AttachPluginToApiRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// API编号
	ApiId string `json:"api_id"`

	Body *ApiOperPluginInfo `json:"body,omitempty"`
}

func (o AttachPluginToApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachPluginToApiRequest struct{}"
	}

	return strings.Join([]string{"AttachPluginToApiRequest", string(data)}, " ")
}
