package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseRouterAttachmentId 企业路由器的连接ID。
type EnterpriseRouterAttachmentId struct {

	// 企业路由器的连接ID。
	EnterpriseRouterAttachmentId *string `json:"enterprise_router_attachment_id,omitempty"`
}

func (o EnterpriseRouterAttachmentId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseRouterAttachmentId struct{}"
	}

	return strings.Join([]string{"EnterpriseRouterAttachmentId", string(data)}, " ")
}
