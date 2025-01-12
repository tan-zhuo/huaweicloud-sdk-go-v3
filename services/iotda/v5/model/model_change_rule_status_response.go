package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeRuleStatusResponse Response Object
type ChangeRuleStatusResponse struct {

	// **参数说明**：规则的激活状态。 **取值范围**： - active：激活。 - inactive：未激活。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeRuleStatusResponse", string(data)}, " ")
}
