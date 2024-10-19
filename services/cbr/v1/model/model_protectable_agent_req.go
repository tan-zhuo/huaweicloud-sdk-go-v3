package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectableAgentReq
type ProtectableAgentReq struct {

	// 查询参数列表
	AgentStatus []ProtectableAgentStatusResource `json:"agent_status"`
}

func (o ProtectableAgentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableAgentReq struct{}"
	}

	return strings.Join([]string{"ProtectableAgentReq", string(data)}, " ")
}
