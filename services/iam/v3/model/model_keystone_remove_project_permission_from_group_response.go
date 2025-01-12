package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneRemoveProjectPermissionFromGroupResponse Response Object
type KeystoneRemoveProjectPermissionFromGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneRemoveProjectPermissionFromGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneRemoveProjectPermissionFromGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveProjectPermissionFromGroupResponse", string(data)}, " ")
}
