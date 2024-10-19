package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDocWatermarkRequest Request Object
type CreateDocWatermarkRequest struct {
	Body *CreateDocWatermarkRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateDocWatermarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocWatermarkRequest struct{}"
	}

	return strings.Join([]string{"CreateDocWatermarkRequest", string(data)}, " ")
}
