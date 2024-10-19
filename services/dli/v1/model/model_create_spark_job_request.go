package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSparkJobRequest Request Object
type CreateSparkJobRequest struct {
	UserId *string `json:"USER-ID,omitempty"`

	Body *CreateSparkJobRequestBody `json:"body,omitempty"`
}

func (o CreateSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJobRequest struct{}"
	}

	return strings.Join([]string{"CreateSparkJobRequest", string(data)}, " ")
}
