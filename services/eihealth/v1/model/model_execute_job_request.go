package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteJobRequest Request Object
type ExecuteJobRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *JobDto `json:"body,omitempty"`
}

func (o ExecuteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteJobRequest struct{}"
	}

	return strings.Join([]string{"ExecuteJobRequest", string(data)}, " ")
}
