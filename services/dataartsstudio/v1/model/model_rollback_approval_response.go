package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackApprovalResponse Response Object
type RollbackApprovalResponse struct {
	Data           *RollbackApprovalResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o RollbackApprovalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackApprovalResponse struct{}"
	}

	return strings.Join([]string{"RollbackApprovalResponse", string(data)}, " ")
}
