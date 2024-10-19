package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachPluginFromApiRequest Request Object
type DetachPluginFromApiRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// API编号
	ApiId string `json:"api_id"`

	Body *ApiOperPluginInfo `json:"body,omitempty"`
}

func (o DetachPluginFromApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachPluginFromApiRequest struct{}"
	}

	return strings.Join([]string{"DetachPluginFromApiRequest", string(data)}, " ")
}
