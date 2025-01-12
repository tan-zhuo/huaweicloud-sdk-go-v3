package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveProjectPermissionFromAgencyResponse Response Object
type RemoveProjectPermissionFromAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveProjectPermissionFromAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveProjectPermissionFromAgencyResponse struct{}"
	}

	return strings.Join([]string{"RemoveProjectPermissionFromAgencyResponse", string(data)}, " ")
}
