package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnAccessPolicyRequestBody struct {
	AccessPolicy *UpdateVpnAccessPolicyRequestBodyContent `json:"access_policy"`
}

func (o UpdateVpnAccessPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnAccessPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpnAccessPolicyRequestBody", string(data)}, " ")
}
