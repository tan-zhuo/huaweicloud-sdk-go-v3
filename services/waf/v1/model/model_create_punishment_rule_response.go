package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePunishmentRuleResponse Response Object
type CreatePunishmentRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 所属策略id
	Policyid *string `json:"policyid,omitempty"`

	// 拦截时间
	BlockTime *int32 `json:"block_time,omitempty"`

	// 攻击惩罚类别
	Category *string `json:"category,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 创建规则时间戳
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreatePunishmentRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePunishmentRuleResponse struct{}"
	}

	return strings.Join([]string{"CreatePunishmentRuleResponse", string(data)}, " ")
}
