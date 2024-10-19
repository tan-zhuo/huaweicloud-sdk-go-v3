package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnterpriseProjectsRequest Request Object
type ShowEnterpriseProjectsRequest struct {
}

func (o ShowEnterpriseProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectsRequest struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectsRequest", string(data)}, " ")
}
