package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCheckResultRequest Request Object
type RunCheckResultRequest struct {

	// 任务标识。
	JobId string `json:"job_id"`
}

func (o RunCheckResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCheckResultRequest struct{}"
	}

	return strings.Join([]string{"RunCheckResultRequest", string(data)}, " ")
}
