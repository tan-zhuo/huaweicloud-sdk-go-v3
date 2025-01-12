package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowRequest Request Object
type CreateWorkflowRequest struct {
	Body *WorkflowCreateBody `json:"body,omitempty"`
}

func (o CreateWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowRequest", string(data)}, " ")
}
