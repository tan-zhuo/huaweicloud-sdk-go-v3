package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseWaterMarkRequest Request Object
type ShowDatabaseWaterMarkRequest struct {
	Body *ExtractedDatabaseWatermark `json:"body,omitempty"`
}

func (o ShowDatabaseWaterMarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseWaterMarkRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseWaterMarkRequest", string(data)}, " ")
}
