package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationRequest Request Object
type DeleteOrganizationRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o DeleteOrganizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationRequest struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationRequest", string(data)}, " ")
}
