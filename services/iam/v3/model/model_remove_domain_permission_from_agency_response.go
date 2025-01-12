package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveDomainPermissionFromAgencyResponse Response Object
type RemoveDomainPermissionFromAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveDomainPermissionFromAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveDomainPermissionFromAgencyResponse struct{}"
	}

	return strings.Join([]string{"RemoveDomainPermissionFromAgencyResponse", string(data)}, " ")
}
