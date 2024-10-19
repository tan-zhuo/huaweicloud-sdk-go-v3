package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCheckDomainPermissionForGroupResponse Response Object
type KeystoneCheckDomainPermissionForGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneCheckDomainPermissionForGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCheckDomainPermissionForGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCheckDomainPermissionForGroupResponse", string(data)}, " ")
}
