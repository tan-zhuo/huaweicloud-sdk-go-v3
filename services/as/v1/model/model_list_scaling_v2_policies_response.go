package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScalingV2PoliciesResponse Response Object
type ListScalingV2PoliciesResponse struct {

	// 总记录数。
	TotalNumber *int32 `json:"total_number,omitempty"`

	// 查询的起始行号。
	StartNumber *int32 `json:"start_number,omitempty"`

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty"`

	// 伸缩策略列表
	ScalingPolicies *[]ScalingPoliciesV2 `json:"scaling_policies,omitempty"`
	HttpStatusCode  int                  `json:"-"`
}

func (o ListScalingV2PoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingV2PoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListScalingV2PoliciesResponse", string(data)}, " ")
}
