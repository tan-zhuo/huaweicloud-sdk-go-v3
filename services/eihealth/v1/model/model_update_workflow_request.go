package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowRequest Request Object
type UpdateWorkflowRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 流程id
	WorkflowId string `json:"workflow_id"`

	Body *WorkflowDto `json:"body,omitempty"`
}

func (o UpdateWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowRequest", string(data)}, " ")
}
