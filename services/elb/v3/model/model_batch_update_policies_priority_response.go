package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePoliciesPriorityResponse Response Object
type BatchUpdatePoliciesPriorityResponse struct {

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdatePoliciesPriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityResponse", string(data)}, " ")
}
