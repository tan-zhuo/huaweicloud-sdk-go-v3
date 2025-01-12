package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCaseRequestBody CreateCaseRequestBody
type CreateCaseRequestBody struct {

	// 名称
	Name string `json:"name"`

	// type（0-常规用例，1-视频流用例，2-预制用例）
	Type int32 `json:"type"`

	// 所属任务id
	TaskId int32 `json:"task_id"`
}

func (o CreateCaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCaseRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCaseRequestBody", string(data)}, " ")
}
