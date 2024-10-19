package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAuditDatabaseResponse Response Object
type SwitchAuditDatabaseResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchAuditDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAuditDatabaseResponse struct{}"
	}

	return strings.Join([]string{"SwitchAuditDatabaseResponse", string(data)}, " ")
}
