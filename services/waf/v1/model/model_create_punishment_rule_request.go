package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePunishmentRuleRequest Request Object
type CreatePunishmentRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	Body *CreatePunishmentRuleRequestBody `json:"body,omitempty"`
}

func (o CreatePunishmentRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePunishmentRuleRequest struct{}"
	}

	return strings.Join([]string{"CreatePunishmentRuleRequest", string(data)}, " ")
}
