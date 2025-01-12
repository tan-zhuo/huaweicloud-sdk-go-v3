package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookTopologyResponse Response Object
type ShowPlaybookTopologyResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 流程实例列表
	ActionInstances *[]ActionInstanceInfo `json:"action_instances,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPlaybookTopologyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookTopologyResponse struct{}"
	}

	return strings.Join([]string{"ShowPlaybookTopologyResponse", string(data)}, " ")
}
