package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateL7ruleRequestBody This is a auto create Body Object
type CreateL7ruleRequestBody struct {
	Rule *CreateL7ruleReq `json:"rule"`
}

func (o CreateL7ruleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7ruleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7ruleRequestBody", string(data)}, " ")
}
