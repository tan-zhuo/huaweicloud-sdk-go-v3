package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneRemoveDomainPermissionFromGroupResponse Response Object
type KeystoneRemoveDomainPermissionFromGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneRemoveDomainPermissionFromGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneRemoveDomainPermissionFromGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveDomainPermissionFromGroupResponse", string(data)}, " ")
}
