package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneAssociateGroupWithProjectPermissionResponse Response Object
type KeystoneAssociateGroupWithProjectPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneAssociateGroupWithProjectPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneAssociateGroupWithProjectPermissionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneAssociateGroupWithProjectPermissionResponse", string(data)}, " ")
}
