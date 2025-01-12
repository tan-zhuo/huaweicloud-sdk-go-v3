package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRemuxTaskResponse Response Object
type CreateRemuxTaskResponse struct {

	// 任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRemuxTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateRemuxTaskResponse", string(data)}, " ")
}
