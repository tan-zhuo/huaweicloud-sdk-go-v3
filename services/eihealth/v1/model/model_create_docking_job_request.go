package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDockingJobRequest Request Object
type CreateDockingJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateDockJobReq `json:"body,omitempty"`
}

func (o CreateDockingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDockingJobRequest struct{}"
	}

	return strings.Join([]string{"CreateDockingJobRequest", string(data)}, " ")
}
