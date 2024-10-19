package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchEnableAlarmPoliciesRequestBody struct {

	// 需要批量启停的告警规则策略的ID列表
	AlarmPolicyIds []string `json:"alarm_policy_ids"`

	// 开关
	Enabled bool `json:"enabled"`
}

func (o BatchEnableAlarmPoliciesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAlarmPoliciesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchEnableAlarmPoliciesRequestBody", string(data)}, " ")
}
