package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEditingJobRequest Request Object
type CreateEditingJobRequest struct {
	Body *CreateEditingJobReq `json:"body,omitempty"`
}

func (o CreateEditingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobRequest struct{}"
	}

	return strings.Join([]string{"CreateEditingJobRequest", string(data)}, " ")
}
