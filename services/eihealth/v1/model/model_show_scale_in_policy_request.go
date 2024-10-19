package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScaleInPolicyRequest Request Object
type ShowScaleInPolicyRequest struct {
}

func (o ShowScaleInPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScaleInPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScaleInPolicyRequest", string(data)}, " ")
}
