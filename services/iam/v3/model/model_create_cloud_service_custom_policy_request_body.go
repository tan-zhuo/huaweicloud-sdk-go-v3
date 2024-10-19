package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudServiceCustomPolicyRequestBody
type CreateCloudServiceCustomPolicyRequestBody struct {
	Role *ServicePolicyRoleOption `json:"role"`
}

func (o CreateCloudServiceCustomPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudServiceCustomPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCloudServiceCustomPolicyRequestBody", string(data)}, " ")
}
