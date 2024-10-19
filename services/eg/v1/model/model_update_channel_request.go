package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateChannelRequest Request Object
type UpdateChannelRequest struct {

	// 指定查询的事件通道ID
	ChannelId string `json:"channel_id"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ChannelUpdateReq `json:"body,omitempty"`
}

func (o UpdateChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChannelRequest struct{}"
	}

	return strings.Join([]string{"UpdateChannelRequest", string(data)}, " ")
}
