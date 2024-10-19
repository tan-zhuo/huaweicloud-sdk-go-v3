package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckAllProjectsPermissionForAgencyResponse Response Object
type CheckAllProjectsPermissionForAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckAllProjectsPermissionForAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAllProjectsPermissionForAgencyResponse struct{}"
	}

	return strings.Join([]string{"CheckAllProjectsPermissionForAgencyResponse", string(data)}, " ")
}
