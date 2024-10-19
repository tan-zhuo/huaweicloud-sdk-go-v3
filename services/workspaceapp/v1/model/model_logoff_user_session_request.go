package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoffUserSessionRequest Request Object
type LogoffUserSessionRequest struct {
	Body *LogoffUserSessionReq `json:"body,omitempty"`
}

func (o LogoffUserSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffUserSessionRequest struct{}"
	}

	return strings.Join([]string{"LogoffUserSessionRequest", string(data)}, " ")
}
