package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeProxyScaleResponse Response Object
type ChangeProxyScaleResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeProxyScaleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeProxyScaleResponse struct{}"
	}

	return strings.Join([]string{"ChangeProxyScaleResponse", string(data)}, " ")
}
