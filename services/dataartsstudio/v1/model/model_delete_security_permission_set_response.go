package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityPermissionSetResponse Response Object
type DeleteSecurityPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityPermissionSetResponse", string(data)}, " ")
}
