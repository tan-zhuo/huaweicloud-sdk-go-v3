package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmRulesResponse Response Object
type DeleteAlarmRulesResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRulesResponse", string(data)}, " ")
}
