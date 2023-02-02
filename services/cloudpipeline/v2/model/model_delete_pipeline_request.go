package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePipelineRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`

	// 语言类型 中文:zh-cn 英文:en-us，默认en-us
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o DeletePipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineRequest struct{}"
	}

	return strings.Join([]string{"DeletePipelineRequest", string(data)}, " ")
}