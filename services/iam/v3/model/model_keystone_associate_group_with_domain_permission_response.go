package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneAssociateGroupWithDomainPermissionResponse Response Object
type KeystoneAssociateGroupWithDomainPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneAssociateGroupWithDomainPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneAssociateGroupWithDomainPermissionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneAssociateGroupWithDomainPermissionResponse", string(data)}, " ")
}
