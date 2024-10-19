package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnAccessPolicyRequestBody struct {
	AccessPolicy *CreateVpnAccessPolicyRequestBodyContent `json:"access_policy"`
}

func (o CreateVpnAccessPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnAccessPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpnAccessPolicyRequestBody", string(data)}, " ")
}
