package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationConformancePackResponse Response Object
type DeleteOrganizationConformancePackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOrganizationConformancePackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationConformancePackResponse struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationConformancePackResponse", string(data)}, " ")
}
