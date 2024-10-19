package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCallbackWorkflowResponse Response Object
type CreateCallbackWorkflowResponse struct {

	// 错误码
	ExecutionId    *string `json:"execution_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCallbackWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCallbackWorkflowResponse struct{}"
	}

	return strings.Join([]string{"CreateCallbackWorkflowResponse", string(data)}, " ")
}
