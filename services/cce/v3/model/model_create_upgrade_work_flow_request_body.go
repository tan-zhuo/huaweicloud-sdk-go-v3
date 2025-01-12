package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUpgradeWorkFlowRequestBody struct {

	// API类型，固定值“WorkFlowTask”，该值不可修改。
	Kind string `json:"kind"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion string `json:"apiVersion"`

	Spec *WorkFlowSpec `json:"spec"`
}

func (o CreateUpgradeWorkFlowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUpgradeWorkFlowRequestBody struct{}"
	}

	return strings.Join([]string{"CreateUpgradeWorkFlowRequestBody", string(data)}, " ")
}
