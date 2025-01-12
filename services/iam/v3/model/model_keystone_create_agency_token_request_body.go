package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateAgencyTokenRequestBody
type KeystoneCreateAgencyTokenRequestBody struct {
	Auth *AgencyTokenAuth `json:"auth"`
}

func (o KeystoneCreateAgencyTokenRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateAgencyTokenRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateAgencyTokenRequestBody", string(data)}, " ")
}
