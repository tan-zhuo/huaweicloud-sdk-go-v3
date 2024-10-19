package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantAccessInfoResponse Response Object
type UpdateTenantAccessInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTenantAccessInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantAccessInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantAccessInfoResponse", string(data)}, " ")
}
