package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportWorkflowRequest Request Object
type ImportWorkflowRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *ImportWorkflowReq `json:"body,omitempty"`
}

func (o ImportWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ImportWorkflowRequest", string(data)}, " ")
}
