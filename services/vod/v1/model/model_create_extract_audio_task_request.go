package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExtractAudioTaskRequest Request Object
type CreateExtractAudioTaskRequest struct {
	Body *ExtractAudioTaskReq `json:"body,omitempty"`
}

func (o CreateExtractAudioTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExtractAudioTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateExtractAudioTaskRequest", string(data)}, " ")
}
