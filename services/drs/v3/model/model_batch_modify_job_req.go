package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyJobReq 批量修改任务请求体
type BatchModifyJobReq struct {

	// 修改任务请求体
	Jobs []ModifyJobReq `json:"jobs"`
}

func (o BatchModifyJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyJobReq struct{}"
	}

	return strings.Join([]string{"BatchModifyJobReq", string(data)}, " ")
}
