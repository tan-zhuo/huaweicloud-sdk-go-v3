package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunGetFileTranslationResultRequest Request Object
type RunGetFileTranslationResultRequest struct {

	// 文档翻译任务标识符。通过文档翻译接口获取。
	JobId string `json:"job_id"`
}

func (o RunGetFileTranslationResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunGetFileTranslationResultRequest struct{}"
	}

	return strings.Join([]string{"RunGetFileTranslationResultRequest", string(data)}, " ")
}
