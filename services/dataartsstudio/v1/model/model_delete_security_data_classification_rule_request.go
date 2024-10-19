package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityDataClassificationRuleRequest Request Object
type DeleteSecurityDataClassificationRuleRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 需要删除的规则id
	Id string `json:"id"`
}

func (o DeleteSecurityDataClassificationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityDataClassificationRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityDataClassificationRuleRequest", string(data)}, " ")
}
