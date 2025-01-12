package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTablePreviewRequest Request Object
type ShowTablePreviewRequest struct {

	// 表ID。
	TableId string `json:"table_id"`
}

func (o ShowTablePreviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTablePreviewRequest struct{}"
	}

	return strings.Join([]string{"ShowTablePreviewRequest", string(data)}, " ")
}
