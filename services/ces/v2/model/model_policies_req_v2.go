package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoliciesReqV2 struct {

	// 策略信息
	Policies []UpdatePolicy `json:"policies"`
}

func (o PoliciesReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesReqV2 struct{}"
	}

	return strings.Join([]string{"PoliciesReqV2", string(data)}, " ")
}
