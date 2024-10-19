package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchApproveApplyResponse Response Object
type BatchApproveApplyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchApproveApplyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchApproveApplyResponse struct{}"
	}

	return strings.Join([]string{"BatchApproveApplyResponse", string(data)}, " ")
}
