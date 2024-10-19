package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGovernancePolicyResponse Response Object
type UpdateGovernancePolicyResponse struct {

	// 结果信息
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGovernancePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGovernancePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateGovernancePolicyResponse", string(data)}, " ")
}
