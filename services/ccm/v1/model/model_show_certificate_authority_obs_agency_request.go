package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateAuthorityObsAgencyRequest Request Object
type ShowCertificateAuthorityObsAgencyRequest struct {
}

func (o ShowCertificateAuthorityObsAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateAuthorityObsAgencyRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateAuthorityObsAgencyRequest", string(data)}, " ")
}
