package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableOrganizationShareRequest Request Object
type DisableOrganizationShareRequest struct {
}

func (o DisableOrganizationShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableOrganizationShareRequest struct{}"
	}

	return strings.Join([]string{"DisableOrganizationShareRequest", string(data)}, " ")
}
