package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTaskRulesetItem struct {

	// 规则集语言
	Language string `json:"language"`

	// 规则集ID,通过调用ListTaskRuleset接口，根据响应参数中的template_id获得
	RuleSetId string `json:"rule_set_id"`

	// 任务语言和规则集的关系是否启用，1是启用，0是未启用
	IfUse string `json:"if_use"`

	// 新/老数据表示，默认1
	Status string `json:"status"`
}

func (o UpdateTaskRulesetItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRulesetItem struct{}"
	}

	return strings.Join([]string{"UpdateTaskRulesetItem", string(data)}, " ")
}
