package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmRulePoliciesResponse Response Object
type UpdateAlarmRulePoliciesResponse struct {

	// 策略信息
	Policies       *[]UpdatePolicy `json:"policies,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateAlarmRulePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRulePoliciesResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRulePoliciesResponse", string(data)}, " ")
}
