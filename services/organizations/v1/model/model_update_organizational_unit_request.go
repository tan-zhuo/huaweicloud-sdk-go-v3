package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationalUnitRequest Request Object
type UpdateOrganizationalUnitRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 与组织单元关联的唯一标识符（ID）。
	OrganizationalUnitId string `json:"organizational_unit_id"`

	Body *UpdateOrganizationalUnitReqBody `json:"body,omitempty"`
}

func (o UpdateOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationalUnitRequest", string(data)}, " ")
}
