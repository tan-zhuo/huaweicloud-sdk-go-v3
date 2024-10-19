package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNextflowWorkflowRequest Request Object
type ShowNextflowWorkflowRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 流程id
	WorkflowId string `json:"workflow_id"`
}

func (o ShowNextflowWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ShowNextflowWorkflowRequest", string(data)}, " ")
}
