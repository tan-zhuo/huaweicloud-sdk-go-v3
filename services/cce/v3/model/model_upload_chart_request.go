package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadChartRequest Request Object
type UploadChartRequest struct {
	Body *UploadChartRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadChartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadChartRequest struct{}"
	}

	return strings.Join([]string{"UploadChartRequest", string(data)}, " ")
}
