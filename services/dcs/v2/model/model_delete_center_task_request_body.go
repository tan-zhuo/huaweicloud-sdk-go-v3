package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCenterTaskRequestBody 删除任务中心任务
type DeleteCenterTaskRequestBody struct {

	// 删除结果
	Message *string `json:"message,omitempty"`
}

func (o DeleteCenterTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCenterTaskRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteCenterTaskRequestBody", string(data)}, " ")
}
