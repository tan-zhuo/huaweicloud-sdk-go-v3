package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAgencyWithAllProjectsPermissionResponse Response Object
type AssociateAgencyWithAllProjectsPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateAgencyWithAllProjectsPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAgencyWithAllProjectsPermissionResponse struct{}"
	}

	return strings.Join([]string{"AssociateAgencyWithAllProjectsPermissionResponse", string(data)}, " ")
}
