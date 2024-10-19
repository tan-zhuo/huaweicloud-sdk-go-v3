package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityDataClassificationRuleRequest Request Object
type UpdateSecurityDataClassificationRuleRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 识别规则id
	Id string `json:"id"`

	Body *DataClassificationRuleOperateDto `json:"body,omitempty"`
}

func (o UpdateSecurityDataClassificationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityDataClassificationRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityDataClassificationRuleRequest", string(data)}, " ")
}
