package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVariableResultJson struct {

	// 变量id
	VariableId *int32 `json:"variable_id,omitempty"`
}

func (o CreateVariableResultJson) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVariableResultJson struct{}"
	}

	return strings.Join([]string{"CreateVariableResultJson", string(data)}, " ")
}
