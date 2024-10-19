package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectListResponse Response Object
type ShowProjectListResponse struct {

	// 结果状态
	Status *string `json:"status,omitempty"`

	// 请求id
	TraceId *string `json:"trace_id,omitempty"`

	// 请求返回结果，接口不同，返回不同
	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowProjectListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectListResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectListResponse", string(data)}, " ")
}
