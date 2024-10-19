package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBatchJobRequest Request Object
type CreateBatchJobRequest struct {
	Body *Job `json:"body,omitempty"`
}

func (o CreateBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchJobRequest struct{}"
	}

	return strings.Join([]string{"CreateBatchJobRequest", string(data)}, " ")
}
