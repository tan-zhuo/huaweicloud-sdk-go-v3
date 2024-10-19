package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Assignment 授权描述。
type Assignment struct {

	// 策略ID。
	PolicyStatementId string `json:"policy_statement_id"`

	// 目标。
	Attach string `json:"attach"`

	AttachType *AttachType `json:"attach_type"`
}

func (o Assignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Assignment struct{}"
	}

	return strings.Join([]string{"Assignment", string(data)}, " ")
}
