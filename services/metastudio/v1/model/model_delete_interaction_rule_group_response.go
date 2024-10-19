package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInteractionRuleGroupResponse Response Object
type DeleteInteractionRuleGroupResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInteractionRuleGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInteractionRuleGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteInteractionRuleGroupResponse", string(data)}, " ")
}
