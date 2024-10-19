package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestTestCasePageParam struct {
	Params *TestCasePageParam `json:"params"`
}

func (o CommRequestTestCasePageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestTestCasePageParam struct{}"
	}

	return strings.Join([]string{"CommRequestTestCasePageParam", string(data)}, " ")
}
