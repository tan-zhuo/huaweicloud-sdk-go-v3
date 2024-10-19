package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStrategyRequest Request Object
type ListStrategyRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0
	Offset int32 `json:"offset"`

	// 每页显示的条目数量
	Limit int32 `json:"limit"`

	// 是否包含租户级规则
	IncludeTenantRuleSet bool `json:"include_tenant_rule_set"`

	// 策略名称，用于模糊查询
	Name *string `json:"name,omitempty"`

	// 是否有效
	IsValid *bool `json:"is_valid,omitempty"`

	// 策略类型
	Type *string `json:"type,omitempty"`

	// 项目ID
	CloudProjectId *string `json:"cloud_project_id,omitempty"`
}

func (o ListStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStrategyRequest struct{}"
	}

	return strings.Join([]string{"ListStrategyRequest", string(data)}, " ")
}
