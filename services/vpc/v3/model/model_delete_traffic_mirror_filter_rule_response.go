package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrafficMirrorFilterRuleResponse Response Object
type DeleteTrafficMirrorFilterRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTrafficMirrorFilterRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrafficMirrorFilterRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrafficMirrorFilterRuleResponse", string(data)}, " ")
}
