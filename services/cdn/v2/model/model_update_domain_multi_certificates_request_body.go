package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDomainMultiCertificatesRequestBody struct {
	Https *UpdateDomainMultiCertificatesRequestBodyContent `json:"https,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequestBody", string(data)}, " ")
}
