package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordParam 录制参数
type RecordParam struct {

	// 录制规则id。
	RecordRuleId string `json:"record_rule_id"`
}

func (o RecordParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordParam struct{}"
	}

	return strings.Join([]string{"RecordParam", string(data)}, " ")
}
