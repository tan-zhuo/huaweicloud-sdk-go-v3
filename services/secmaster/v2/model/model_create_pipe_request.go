package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipeRequest Request Object
type CreatePipeRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *CreatePipeRequestBody `json:"body,omitempty"`
}

func (o CreatePipeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipeRequest struct{}"
	}

	return strings.Join([]string{"CreatePipeRequest", string(data)}, " ")
}
