package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectTranscriberJobRequest Request Object
type CollectTranscriberJobRequest struct {

	// 录音文件识别任务标识符。
	JobId string `json:"job_id"`
}

func (o CollectTranscriberJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectTranscriberJobRequest struct{}"
	}

	return strings.Join([]string{"CollectTranscriberJobRequest", string(data)}, " ")
}
