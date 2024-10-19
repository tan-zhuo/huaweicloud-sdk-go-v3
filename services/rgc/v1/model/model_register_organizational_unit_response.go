package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterOrganizationalUnitResponse Response Object
type RegisterOrganizationalUnitResponse struct {

	// 异步接口的操作ID。
	OrganizationalUnitOperationId *string `json:"organizational_unit_operation_id,omitempty"`
	HttpStatusCode                int     `json:"-"`
}

func (o RegisterOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"RegisterOrganizationalUnitResponse", string(data)}, " ")
}
