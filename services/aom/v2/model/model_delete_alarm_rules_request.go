package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmRulesRequest Request Object
type DeleteAlarmRulesRequest struct {
	Body *DeleteAlarmRulesBody `json:"body,omitempty"`
}

func (o DeleteAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRulesRequest", string(data)}, " ")
}
