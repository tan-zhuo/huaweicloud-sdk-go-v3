package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateResourceSharePermissionResponse Response Object
type AssociateResourceSharePermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateResourceSharePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResourceSharePermissionResponse struct{}"
	}

	return strings.Join([]string{"AssociateResourceSharePermissionResponse", string(data)}, " ")
}
