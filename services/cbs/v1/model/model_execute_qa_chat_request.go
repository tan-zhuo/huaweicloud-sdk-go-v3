package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteQaChatRequest Request Object
type ExecuteQaChatRequest struct {

	// 机器人标识符，qabot编号，UUID格式。
	QabotId string `json:"qabot_id"`

	Body *PostRequestsReq `json:"body,omitempty"`
}

func (o ExecuteQaChatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteQaChatRequest struct{}"
	}

	return strings.Join([]string{"ExecuteQaChatRequest", string(data)}, " ")
}
