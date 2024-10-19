package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCbhResponse Response Object
type CreateCbhResponse struct {

	// 云堡垒机实例key。
	InstanceKey *int32 `json:"instance_key,omitempty"`

	// 云堡垒机备机实例key。（当前字段未启用，返回默认值null）
	SlaveInstanceKey *int32 `json:"slave_instance_key,omitempty"`

	// 返回创建云堡垒机实例信息。
	RequestInfo *int32 `json:"request_info,omitempty"`

	// job任务ID。（当前字段未启用，返回默认值null）
	JobId          *int32 `json:"job_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateCbhResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCbhResponse struct{}"
	}

	return strings.Join([]string{"CreateCbhResponse", string(data)}, " ")
}
